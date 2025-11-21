package handler

import (
	"encoding/json"
	"net/http"
)

func (h *HandlerImpl) badRequestResponse(w http.ResponseWriter, status int, code ErrorResponseErrorCode, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	errResp := ErrorResponse{
		Error: struct {
			Code    ErrorResponseErrorCode "json:\"code\""
			Message string                 "json:\"message\""
		}{
			Code:    code,
			Message: message,
		},
	}
	json.NewEncoder(w).Encode(errResp)
}
func (h *HandlerImpl) successRequestResponse(w http.ResponseWriter, status int, response any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}
