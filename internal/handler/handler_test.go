package handler_test

// import (
// 	"bytes"
// 	"net/http"
// 	"net/http/httptest"
// 	"os"
// 	"testing"

// 	"github.com/crema-labs/sxg-go/internal/handler"
// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// 	"go.uber.org/zap"
// )

// func TestHandlePost(t *testing.T) {
// 	// Set Gin to Test Mode
// 	gin.SetMode(gin.TestMode)

// 	// Test cases
// 	tests := []struct {
// 		name           string
// 		inputJSON      string
// 		expectedStatus int
// 		expectedBody   string
// 	}{
// 		// {
// 		// 	name:           "Valid input",
// 		// 	inputJSON:      `{}`, // TODO: Fill in with valid input
// 		// 	expectedStatus: http.StatusOK,
// 		// 	expectedBody:   `{"message":"POST request received"}`,
// 		// },
// 		{
// 			name:           "Valid Data",
// 			inputJSON:      `{"source_url":"https://blog.crema.sh","data":"Building foundational libraries to remove limitations for the future projects and "}`,
// 			expectedStatus: http.StatusBadRequest,
// 			expectedBody:   `{"reqId":"bb313c35814c7cbd828c875c29cc4af3cc07e90a770e3a4bbf79af7b4937918e"}`,
// 		},
// 		// {
// 		// 	name:           "Invalid input - empty source_url",
// 		// 	inputJSON:      `{"data":"some data"}`,
// 		// 	expectedStatus: http.StatusBadRequest,
// 		// 	expectedBody:   `{"error":"SourceUrl is required"}`,
// 		// },
// 		// TODO: Add more test cases as needed
// 	}

// 	logger, _ := zap.NewDevelopment()

// 	priv_key, ok := os.LookupEnv("SP1_PRIVATE_KEY")
// 	if !ok {
// 		t.Fatal("SP1_PRIVATE_KEY environment variable is required")
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			// Create a response recorder
// 			w := httptest.NewRecorder()

// 			// Create a request
// 			req, _ := http.NewRequest("POST", "/proof", bytes.NewBufferString(tt.inputJSON))
// 			req.Header.Set("Content-Type", "application/json")

// 			// Create a Gin router with the handler
// 			hpr := handler.HandleProofRequest{
// 				Logger:  logger,
// 				PrivKey: priv_key,
// 			}
// 			r := gin.New()
// 			r.GET("/status", hpr.GetProofStatus)

// 			// Serve the request
// 			r.ServeHTTP(w, req)

// 			// Assert the status code
// 			assert.Equal(t, tt.expectedStatus, w.Code)

// 			// Assert the response body
// 			assert.JSONEq(t, tt.expectedBody, w.Body.String())
// 		})
// 	}
// }

// func TestHandleGet(t *testing.T) {
// 	// Set Gin to Test Mode
// 	gin.SetMode(gin.TestMode)

// 	// Test cases
// 	tests := []struct {
// 		name           string
// 		query          string
// 		expectedStatus int
// 		expectedBody   string
// 	}{
// 		{
// 			name:           "Valid input",
// 			query:          "reqId=bb313c35814c7cbd828c875c29cc4af3cc07e90a770e3a4bbf79af7b4937918e",
// 			expectedStatus: http.StatusOK,
// 			expectedBody:   `{"status":"completed","result":{"proof":"proof","data":"data"}}`,
// 		},
// 		{
// 			name:           "Invalid input - reqId not found",
// 			query:          "reqId=invalid",
// 			expectedStatus: http.StatusNotFound,
// 			expectedBody:   `{"error":"proof request not found"}`,
// 		},
// 	}

// 	logger, _ := zap.NewDevelopment()

// 	priv_key, ok := os.LookupEnv("SP1_PRIVATE_KEY")
// 	if !ok {
// 		t.Fatal("SP1_PRIVATE_KEY environment variable is required")
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			// Create a response recorder
// 			w := httptest.NewRecorder()

// 			// Create a request
// 			req, _ := http.NewRequest("GET", "/status?"+tt.query, nil)

// 			// Create a Gin router with the handler
// 			hpr := handler.HandleProofRequest{
// 				Logger:  logger,
// 				PrivKey: priv_key,
// 			}
// 			r := gin.New()
// 			r.GET("/status", hpr.GetProofStatus)

// 			// Serve the request
// 			r.ServeHTTP(w, req)

// 			// Assert the status code
// 			assert.Equal(t, tt.expectedStatus, w.Code)

// 			// Assert the response body
// 			assert.JSONEq(t, tt.expectedBody, w.Body.String())
// 		})
// 	}
// }
