package handler

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
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
	h.logger.ZapLogger.Info("An bad response has been sent", zap.Int("status", status), zap.Any("code", code), zap.String("message", message))
}
func (h *HandlerImpl) successRequestResponse(w http.ResponseWriter, status int, response any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
	h.logger.ZapLogger.Info("An success response has been sent", zap.Int("status", status), zap.Any("response", response))
}
