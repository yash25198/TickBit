package handler

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/WICG/webpackage/go/signedexchange"
	"github.com/WICG/webpackage/go/signedexchange/version"
	"github.com/crema-labs/sxg-go/pkg/ethereum"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var latestVersion = string(version.AllVersions[len(version.AllVersions)-1])

var ErrDatNotFound = fmt.Errorf("Data not found in the signed exchange")

type ProofRequest struct {
	BlockNumber uint64 `json:"block_number"`
	BlockHash   string `json:"block_hash"`
}

type StatusType string

const (
	Processing StatusType = "Processing"
	Ready      StatusType = "Ready"
)

type StatusResponse struct {
	Status     StatusType     `json:"status,omitempty"`
	TrackingId string         `json:"tracking_id,omitempty"`
	Proof      *ProofResponse `json:"proof,omitempty"`
}

type ProofResponse struct {
	Result          int    `json:"result"`
	VerificationKey string `json:"vkey"`
	PublicKey       string `json:"publicValues"`
	ProofBytes      string `json:"proof"`
}

type ProofInput struct {
	FinalPayload        []byte      `json:"final_payload"`
	IntegrityStartIndex int         `json:"integrity_start_index"`
	Payload             []byte      `json:"payload"`
	R                   [32]byte    `json:"r"`
	S                   [32]byte    `json:"s"`
	Px                  [32]byte    `json:"px"`
	Py                  [32]byte    `json:"py"`
	BlockParams         BlockParams `json:"block_params"`
}

type BlockParams struct {
	BlockNumber uint64 `json:"block_number"`
	BlockHash   string `json:"block_hash"`
}

func (pi ProofInput) MarshalJSON() ([]byte, error) {
	type Alias ProofInput
	return json.Marshal(&struct {
		FinalPayload        []int       `json:"final_payload"`
		IntegrityStartIndex int         `json:"integrity_start_index"`
		Payload             []int       `json:"payload"`
		R                   [32]byte    `json:"r"`
		S                   [32]byte    `json:"s"`
		Px                  [32]byte    `json:"px"`
		Py                  [32]byte    `json:"py"`
		BlockParams         BlockParams `json:"block_params"`
		*Alias
	}{
		FinalPayload:        bytesToIntSlice(pi.FinalPayload),
		IntegrityStartIndex: pi.IntegrityStartIndex,
		Payload:             bytesToIntSlice(pi.Payload),
		R:                   pi.R,
		S:                   pi.S,
		Px:                  pi.Px,
		Py:                  pi.Py,
		BlockParams:         pi.BlockParams,
	})
}

func bytesToIntSlice(b []byte) []int {
	ints := make([]int, len(b))
	for i, v := range b {
		ints[i] = int(v)
	}
	return ints
}

type HandleProofRequest struct {
	TBClient ethereum.TBClient
	Logger   *zap.Logger
	PrivKey  string
}

func NewHandleProofRequest(logger *zap.Logger, privKey string) *HandleProofRequest {
	return &HandleProofRequest{
		Logger:  logger,
		PrivKey: privKey,
	}
}

func extractTrackingID(logContent string) string {
	lines := strings.Split(logContent, "\n")
	for _, line := range lines {
		if strings.Contains(line, "View in explorer: https://explorer.succinct.xyz/") {
			parts := strings.Split(line, "/")
			if len(parts) > 0 {
				return strings.TrimSpace(parts[len(parts)-1])
			}
		}
	}
	return ""
}

func (hp *HandleProofRequest) HandleDeets(c *gin.Context) {
	blockNumber := c.Query("block_number")
	bn, ok := new(big.Int).SetString(blockNumber, 10)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bn should be valid number",
		})
		return
	}

	bets, err := hp.TBClient.GetBets(c, bn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	pool, err := hp.TBClient.GetPoolValue(c, bn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"bets": bets,
		"pool": pool,
	})

}

func (hp *HandleProofRequest) ProofStatus(blockNumber uint64) (StatusResponse, error) {
	reqId := getReqId(blockNumber)
	res := StatusResponse{}

	logFileName := fmt.Sprintf("%s.log", reqId)
	if _, err := os.Stat(logFileName); err == nil {
		// Log file exists, let's parse it for the tracking ID
		logContent, err := os.ReadFile(logFileName)
		if err == nil {
			trackingID := extractTrackingID(string(logContent))
			if trackingID != "" {
				res.TrackingId = fmt.Sprintf("https://explorer.succinct.xyz/%s", trackingID)
			}
		}
	}

	// Check if reqId-fixture.json exists
	filename := fmt.Sprintf("%s-fixture.json", reqId)
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		if _, err := os.Stat(fmt.Sprintf("%s.log", reqId)); os.IsNotExist(err) {
			return res, fmt.Errorf("proof request not found")
		} else {
			res.Status = Processing
			return res, nil
		}
	}

	// Read the fixture file
	data, err := os.ReadFile(filename)
	if err != nil {
		hp.Logger.Error("Failed to read file", zap.Error(err))
		return res, fmt.Errorf("internal server error")
	}

	proof := ProofResponse{}

	if err := json.Unmarshal(data, &proof); err != nil {
		hp.Logger.Error("Failed to unmarshal JSON", zap.Error(err))
		return res, fmt.Errorf("internal server error")
	}

	res.Status = Ready
	res.Proof = &proof
	return res, nil
}

func (hp *HandleProofRequest) HandleGetStatus(c *gin.Context) {
	reqId, err := strconv.ParseUint(c.Query("block_number"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "block_number is required",
		})
		return
	}

	res, err := hp.ProofStatus(reqId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

func getReqId(blockNumber uint64) string {
	reqid := sha256.Sum256([]byte(fmt.Sprintf("%d", blockNumber)))
	return strings.ToLower(hex.EncodeToString(reqid[:]))
}
func (hp *HandleProofRequest) GenerateProofRequest(req ProofRequest) (string, error) {
	reqidStr := getReqId(req.BlockNumber)

	if err := validateProofRequest(req, reqidStr); err != nil {
		return "", err
	}

	sourceUrl := `https://btc.cryptoid.info/btc/block.dws?` + strconv.FormatUint(req.BlockNumber, 10)

	data := `blockID = %d, blockHash = '%s',`

	data = fmt.Sprintf(data, req.BlockNumber, req.BlockHash)

	sxg, err := verifySXG(sourceUrl)
	if err != nil {
		return "", fmt.Errorf("Failed to verify SXG: %v", err)
	}

	dataStartIndex := findByteArray(convertToBytes(data), sxg.Payload)
	if dataStartIndex <= 0 {
		return "", ErrDatNotFound
	}

	integrity := []byte{}
	enc := sxg.Version.MiceEncoding()
	if d := sxg.ResponseHeaders.Get("Digest"); d != "" {
		integrity = []byte(d)
	} else {
		var buf bytes.Buffer
		d, err := enc.Encode(&buf, sxg.Payload, 16384)
		if err != nil {
			return "", err
		}
		sxg.Payload = buf.Bytes()
		integrity = []byte(d)
	}

	integrityStartIndex := findByteArray(integrity, sxg.Msg)
	if integrityStartIndex <= 0 {
		return "", fmt.Errorf("Integrity not found in the signed message")
	}

	r := [32]byte{}
	copy(r[:], sxg.R.FillBytes(r[:]))

	s := [32]byte{}
	copy(s[:], sxg.S.FillBytes(s[:]))

	px := [32]byte{}
	copy(px[:], sxg.Px.FillBytes(px[:]))

	py := [32]byte{}
	copy(py[:], sxg.Py.FillBytes(py[:]))

	pi := ProofInput{
		FinalPayload:        sxg.Msg,
		IntegrityStartIndex: integrityStartIndex,
		Payload:             sxg.Payload,
		R:                   r,
		S:                   s,
		Px:                  px,
		Py:                  py,
		BlockParams: BlockParams{
			BlockNumber: req.BlockNumber,
			BlockHash:   req.BlockHash,
		},
	}

	// Marshal the ProofInput struct to JSON
	jsonData, err := json.Marshal(pi)
	if err != nil {
		return "", fmt.Errorf("Failed to marshal JSON data: %v", err)
	}

	// Write the JSON data to a file
	filename := fmt.Sprintf("%s.json", reqidStr)
	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return "", fmt.Errorf("Failed to write JSON data to file: %v", err)
	}

	// Create a log file for the command output
	logFileName := fmt.Sprintf("%s.log", reqidStr)
	logFile, err := os.Create(logFileName)
	if err != nil {
		return "", fmt.Errorf("Failed to create log file: %v", err)
	}
	defer logFile.Close()

	// Prepare the CLI command
	cmdArgs := []string{"--system", "groth16", "--input-file-id", reqidStr}
	cmd := exec.Command("/Users/yash/Desktop/crema/sxg-go/cmd/bin/evm", cmdArgs...)

	cmd.Stdout = logFile
	cmd.Stderr = logFile
	cmd.Env = append(os.Environ(), "SP1_PROVER=network", "SP1_PRIVATE_KEY="+hp.PrivKey, "RUST_LOG=info")

	// cmd.SysProcAttr = &syscall.SysProcAttr{
	// 	Setpgid: true,
	// 	Pgid:    0,
	// 	Setsid: true,
	// }

	err = cmd.Start()
	if err != nil {
		hp.Logger.Error("Failed to start Rust program", zap.Error(err))
		return "", fmt.Errorf("Failed to start Rust program: %v", err)
	}

	// Log the process ID for debugging purposes
	hp.Logger.Info("Started Rust program", zap.Int("pid", cmd.Process.Pid))

	return reqidStr, nil
}

func (hp *HandleProofRequest) HandleProofRequest(c *gin.Context) {
	var req ProofRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	reqId, err := hp.GenerateProofRequest(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tracking_id": reqId,
	})
}

func verifySXG(sourceUrl string) (*signedexchange.Exchange, error) {
	client := http.DefaultClient
	req, err := http.NewRequest("GET", sourceUrl, nil)
	if err != nil {
		return nil, err
	}
	ver, ok := version.Parse(latestVersion)
	if !ok {
		return nil, fmt.Errorf("failed to parse version %q", latestVersion)
	}
	mimeType := ver.MimeType()
	req.Header.Add("Accept", mimeType)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	respMimeType := resp.Header.Get("Content-Type")
	if respMimeType != mimeType {
		return nil, fmt.Errorf("GET %q responded with unexpected content type %q", sourceUrl, respMimeType)
	}
	var in io.Reader = resp.Body
	defer resp.Body.Close()

	e, err := signedexchange.ReadExchange(in)
	if err != nil {
		return nil, err
	}

	certFetcher := signedexchange.DefaultCertFetcher

	if err := verify(e, certFetcher, time.Now()); err != nil {
		return nil, err
	}

	return e, nil
}

func verify(e *signedexchange.Exchange, certFetcher signedexchange.CertFetcher, verificationTime time.Time) error {
	if decodedPayload, ok := e.Verify(verificationTime, certFetcher, log.New(os.Stdout, "", 0)); ok {
		e.Payload = decodedPayload
		return nil
	}
	return fmt.Errorf("The exchange has an invalid signature.")
}

func findByteArray(a, b []byte) int {
	if len(a) == 0 {
		return 0
	}
	if len(a) > len(b) {
		return -1
	}

	return bytes.Index(b, a)
}

func validateProofRequest(req ProofRequest, reqId string) error {
	if req.BlockNumber == 0 {
		return fmt.Errorf("BlockNumber is required")
	}

	if req.BlockHash == "" {
		return fmt.Errorf("BlockHash is required")
	}

	// // Check if reqId.json exists
	filename := fmt.Sprintf("%s.json", reqId)
	if _, err := os.Stat(filename); err == nil {
		return fmt.Errorf("proof request already exists, try getting status, reqId: %s", reqId)
	}

	return nil
}

func convertToBytes(s string) []byte {
	var res []byte
	strlen := len(s)
	for i := 0; i < strlen; i++ {
		if i < strlen-1 && s[i] == '\\' && s[i+1] == 'r' {
			res = append(res, 27)
			continue
		}
		res = append(res, byte(s[i]))
	}
	return res
}
