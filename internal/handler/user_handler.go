package handler

import (
	mye "AvitoTest1/internal/errors"
	"encoding/json"
	"net/http"
)

func (h *HandlerImpl) GetUsersGetReview(w http.ResponseWriter, r *http.Request, params GetUsersGetReviewParams) {
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
	var req PostUsersSetIsActiveJSONBody
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.badRequestResponse(w, http.StatusNotFound, NOTFOUND, mye.ErrResourceNotFound.Error())
		return
	}
	user, err := h.usService.SetIsActive(r.Context(), req.UserId, req.IsActive)
	if err != nil {
		h.badRequestResponse(w, http.StatusNotFound, NOTFOUND, err.Error())
		return
	}
	h.successRequestResponse(w, http.StatusOK, user)
}
