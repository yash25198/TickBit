package types

type ProofResponse struct {
	Merkle []string `json:"merkle"`
	Pos    uint64   `json:"pos"`
	Height uint64   `json:"block_height"`
}

type BlockDetails struct {
	BlockHash   string
	Header      []byte
	BlockNumber uint64
}

type Transaction struct {
	TxID   string `json:"txid"`
	VINs   []VIN  `json:"vin"`
	Status Status `json:"status"`
}

type VIN struct {
	TxID         string    `json:"txid"`
	Vout         int       `json:"vout"`
	Prevout      Prevout   `json:"prevout"`
	ScriptSigAsm string    `json:"scriptsig_asm"`
	Witness      *[]string `json:"witness" `
}

type Prevout struct {
	ScriptPubKeyType    string `json:"scriptpubkey_type"`
	ScriptPubKey        string `json:"scriptpubkey"`
	ScriptPubKeyAddress string `json:"scriptpubkey_address"`
}

type Status struct {
	Confirmed   bool   `json:"confirmed"`
	BlockHeight uint64 `json:"block_height"`
}
