package handler

import (
	mye "AvitoTest1/internal/errors"
	"AvitoTest1/internal/models"
	"encoding/json"
	"net/http"
)

func (h *HandlerImpl) PostTeamAdd(w http.ResponseWriter, r *http.Request) {
	var req models.Team
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.badRequestResponse(w, 400, TEAMEXISTS, mye.ErrTeamExist.Error())
		return
	}
	err = h.tService.AddTeam(r.Context(), &req)
	if err != nil {
		h.badRequestResponse(w, 400, TEAMEXISTS, err.Error())
		return
	}
	h.successRequestResponse(w, 201, req)
}

func (h *HandlerImpl) GetTeamGet(w http.ResponseWriter, r *http.Request, params GetTeamGetParams) {
	team, err := h.tService.GetTeam(r.Context(), params.TeamName)
	if err != nil {
		h.badRequestResponse(w, http.StatusNotFound, NOTFOUND, err.Error())
		return
	}
	h.successRequestResponse(w, http.StatusOK, team)
}
