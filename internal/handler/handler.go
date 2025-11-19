package handler

import (
	"AvitoTest1/internal/models"
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

type UserService interface {
	SetIsActive(ctx context.Context, userId string, status bool) (*models.User, error)
}
type HandlerImpl struct {
	usService UserService
}

func (h *HandlerImpl) PostPullRequestCreate(w http.ResponseWriter, r *http.Request) {

}
func (h *HandlerImpl) PostPullRequestMerge(w http.ResponseWriter, r *http.Request) {

}
func (h *HandlerImpl) PostPullRequestReassign(w http.ResponseWriter, r *http.Request) {

}
func (h *HandlerImpl) PostTeamAdd(w http.ResponseWriter, r *http.Request) {

}

func (h *HandlerImpl) GetTeamGet(w http.ResponseWriter, r *http.Request, params GetTeamGetParams) {

}

func (h *HandlerImpl) GetUsersGetReview(w http.ResponseWriter, r *http.Request, params GetUsersGetReviewParams) {

}
func (h *HandlerImpl) PostUsersSetIsActive(w http.ResponseWriter, r *http.Request) {
	var req PostUsersSetIsActiveJSONBody
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		errResp := ErrorResponse{
			Error: struct {
				Code    ErrorResponseErrorCode "json:\"code\""
				Message string                 "json:\"message\""
			}{
				Code:    NOTFOUND,
				Message: "user not found",
			},
		}
		json.NewEncoder(w).Encode(errResp)
	}
	user, err := h.usService.SetIsActive(r.Context(), req.UserId, req.IsActive)
	if err != nil {
		if strings.HasPrefix(err.Error(), "client: ") {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			errResp := ErrorResponse{
				Error: struct {
					Code    ErrorResponseErrorCode "json:\"code\""
					Message string                 "json:\"message\""
				}{
					Code:    NOTFOUND,
					Message: "user not found",
				},
			}
			json.NewEncoder(w).Encode(errResp)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			errResp := ErrorResponse{
				Error: struct {
					Code    ErrorResponseErrorCode "json:\"code\""
					Message string                 "json:\"message\""
				}{
					Code:    INTERNALSERVER,
					Message: "internal server error",
				},
			}
			json.NewEncoder(w).Encode(errResp)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
