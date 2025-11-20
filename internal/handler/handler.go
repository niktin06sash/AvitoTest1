package handler

import (
	"AvitoTest1/internal/models"
	"context"
	"encoding/json"
	"net/http"
)

type UserService interface {
	SetIsActive(ctx context.Context, userId string, status bool) (*models.User, error)
	GetUsersReview(ctx context.Context, userid string) ([]*models.PullRequestShort, error)
}
type TeamService interface {
	AddTeam(ctx context.Context, team *models.Team) error
}
type HandlerImpl struct {
	usService UserService
	tService  TeamService
}

func (h *HandlerImpl) PostPullRequestCreate(w http.ResponseWriter, r *http.Request) {

}
func (h *HandlerImpl) PostPullRequestMerge(w http.ResponseWriter, r *http.Request) {

}
func (h *HandlerImpl) PostPullRequestReassign(w http.ResponseWriter, r *http.Request) {

}
func (h *HandlerImpl) PostTeamAdd(w http.ResponseWriter, r *http.Request) {
	var req models.Team
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		errResp := ErrorResponse{
			Error: struct {
				Code    ErrorResponseErrorCode "json:\"code\""
				Message string                 "json:\"message\""
			}{
				Code:    TEAMEXISTS,
				Message: "team_name already exists",
			},
		}
		json.NewEncoder(w).Encode(errResp)
		return
	}
	err = h.tService.AddTeam(r.Context(), &req)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		errResp := ErrorResponse{
			Error: struct {
				Code    ErrorResponseErrorCode "json:\"code\""
				Message string                 "json:\"message\""
			}{
				Code:    TEAMEXISTS,
				Message: "team_name already exists",
			},
		}
		json.NewEncoder(w).Encode(errResp)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(req)
}

func (h *HandlerImpl) GetTeamGet(w http.ResponseWriter, r *http.Request, params GetTeamGetParams) {

}

func (h *HandlerImpl) GetUsersGetReview(w http.ResponseWriter, r *http.Request, params GetUsersGetReviewParams) {
	pr, err := h.usService.GetUsersReview(r.Context(), params.UserId)
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
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"user_id":       params.UserId,
		"pull_requests": pr,
	}
	json.NewEncoder(w).Encode(response)
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
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
