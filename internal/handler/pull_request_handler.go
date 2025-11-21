package handler

import (
	mye "AvitoTest1/internal/errors"
	"AvitoTest1/internal/models"
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type UserService interface {
	SetIsActive(ctx context.Context, userId string, status bool) (*models.User, error)
	GetUsersReview(ctx context.Context, userid string) ([]*models.PullRequestShort, error)
}
type TeamService interface {
	AddTeam(ctx context.Context, team *models.Team) error
	GetTeam(ctx context.Context, teamname string) (*models.Team, error)
}
type PullRequestService interface {
	CreatePullRequest(ctx context.Context, authorID string, id string, name string) (*models.PullRequest, error)
	MergePullRequest(ctx context.Context, prID string) (*models.PullRequest, error)
	ReassignPullRequest(ctx context.Context, olduserID string, prID string) (*models.PullRequest, string, error)
}
type HandlerImpl struct {
	usService UserService
	tService  TeamService
	prService PullRequestService
}

func (h *HandlerImpl) PostPullRequestCreate(w http.ResponseWriter, r *http.Request) {
	var req PostPullRequestCreateJSONBody
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.badRequestResponse(w, http.StatusNotFound, NOTFOUND, mye.ErrResourceNotFound.Error())
		return
	}
	pr, err := h.prService.CreatePullRequest(r.Context(), req.AuthorId, req.PullRequestId, req.PullRequestName)
	if err != nil {
		if errors.Is(err, mye.ErrResourceNotFound) {
			h.badRequestResponse(w, http.StatusNotFound, NOTFOUND, err.Error())
		} else if errors.Is(err, mye.ErrPRExist) {
			h.badRequestResponse(w, 409, PREXISTS, err.Error())
		}
		return
	}
	h.successRequestResponse(w, 201, pr)
}
func (h *HandlerImpl) PostPullRequestMerge(w http.ResponseWriter, r *http.Request) {
	var req PostPullRequestMergeJSONBody
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.badRequestResponse(w, http.StatusNotFound, NOTFOUND, mye.ErrResourceNotFound.Error())
		return
	}
	pr, err := h.prService.MergePullRequest(r.Context(), req.PullRequestId)
	if err != nil {
		h.badRequestResponse(w, http.StatusNotFound, NOTFOUND, err.Error())
		return
	}
	h.successRequestResponse(w, http.StatusOK, pr)
}
func (h *HandlerImpl) PostPullRequestReassign(w http.ResponseWriter, r *http.Request) {
	var req PostPullRequestReassignJSONBody
	pr, newrew, err := h.prService.ReassignPullRequest(r.Context(), req.OldUserId, req.PullRequestId)
	if err != nil {
		if errors.Is(err, mye.ErrResourceNotFound) {
			h.badRequestResponse(w, http.StatusNotFound, NOTFOUND, err.Error())
		} else if errors.Is(err, mye.ErrReviewerNotAssigned) {
			h.badRequestResponse(w, 409, NOTASSIGNED, err.Error())
		} else if errors.Is(err, mye.ErrNoActiveCandidate) {
			h.badRequestResponse(w, 409, NOCANDIDATE, err.Error())
		} else if errors.Is(err, mye.ErrMergedPR) {
			h.badRequestResponse(w, 409, PRMERGED, err.Error())
		}
		return
	}
	response := map[string]interface{}{
		"pr":          pr,
		"replaced_by": newrew,
	}
	h.successRequestResponse(w, http.StatusOK, response)
}
