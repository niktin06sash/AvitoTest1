package handler

import "net/http"

type HandlerImpl struct {
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

}
