package boundary

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	ErrorCode string `json:"error"`
	Message   string `json:"message"`
}
type SuccessResponse struct {
	ResponseData interface{} `json:"Data"`
	Message      string      `json:"message"`
}

func WriteResponseSuccess(w http.ResponseWriter, statusCode int, successResp SuccessResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(successResp)
}

func WriteResponseErr(w http.ResponseWriter, statusCode int, errResp ErrorResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(errResp)
}
