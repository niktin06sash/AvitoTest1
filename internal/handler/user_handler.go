package handler

import (
	mye "AvitoTest1/internal/errors"
	"AvitoTest1/internal/models"
	"context"
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

type UserService interface {
	SetIsActive(ctx context.Context, userId string, status bool) (*models.User, error)
	GetUsersReview(ctx context.Context, userid string) ([]*models.PullRequestShort, error)
}

func (h *HandlerImpl) GetUsersGetReview(w http.ResponseWriter, r *http.Request, params GetUsersGetReviewParams) {
	h.logger.ZapLogger.Info("New GetUsersGetReview", zap.String("method", r.Method), zap.String("path", r.URL.Path), zap.String("remote_addr", r.RemoteAddr))
	h.logger.ZapLogger.Info("Processing getting user's reviews...", zap.Any("request_body", params))
	pr, err := h.usService.GetUsersReview(r.Context(), params.UserId)
	if err != nil {
		h.badRequestResponse(w, http.StatusNotFound, NOTFOUND, err.Error())
		return
	}
	response := map[string]interface{}{
		"user_id":       params.UserId,
		"pull_requests": pr,
	}
	h.successRequestResponse(w, http.StatusOK, response)
}
func (h *HandlerImpl) PostUsersSetIsActive(w http.ResponseWriter, r *http.Request) {
	h.logger.ZapLogger.Info("New PostUsersSetIsActive", zap.String("method", r.Method), zap.String("path", r.URL.Path), zap.String("remote_addr", r.RemoteAddr))
	var req PostUsersSetIsActiveJSONBody
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.badRequestResponse(w, http.StatusNotFound, NOTFOUND, mye.ErrResourceNotFound.Error())
		return
	}
	h.logger.ZapLogger.Info("Processing setting user's active...", zap.Any("request_body", req))
	user, err := h.usService.SetIsActive(r.Context(), req.UserId, req.IsActive)
	if err != nil {
		h.badRequestResponse(w, http.StatusNotFound, NOTFOUND, err.Error())
		return
	}
	h.successRequestResponse(w, http.StatusOK, user)
}
