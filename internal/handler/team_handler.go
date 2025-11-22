package handler

import (
	mye "AvitoTest1/internal/errors"
	"AvitoTest1/internal/models"
	"context"
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

type TeamService interface {
	AddTeam(ctx context.Context, team *models.Team) error
	GetTeam(ctx context.Context, teamname string) (*models.Team, error)
}

func (h *HandlerImpl) PostTeamAdd(w http.ResponseWriter, r *http.Request) {
	h.logger.ZapLogger.Info("New PostTeamAdd", zap.String("method", r.Method), zap.String("path", r.URL.Path), zap.String("remote_addr", r.RemoteAddr))
	var req models.Team
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.badRequestResponse(w, 400, TEAMEXISTS, mye.ErrTeamExist.Error())
		return
	}
	h.logger.ZapLogger.Info("Processing adding team...", zap.Any("request_body", req))
	err = h.tService.AddTeam(r.Context(), &req)
	if err != nil {
		h.badRequestResponse(w, 400, TEAMEXISTS, err.Error())
		return
	}
	h.successRequestResponse(w, 201, req)
}

func (h *HandlerImpl) GetTeamGet(w http.ResponseWriter, r *http.Request, params GetTeamGetParams) {
	h.logger.ZapLogger.Info("New GetTeamGet", zap.String("method", r.Method), zap.String("path", r.URL.Path), zap.String("remote_addr", r.RemoteAddr))
	h.logger.ZapLogger.Info("Processing getting team...", zap.Any("request_body", params))
	team, err := h.tService.GetTeam(r.Context(), params.TeamName)
	if err != nil {
		h.badRequestResponse(w, http.StatusNotFound, NOTFOUND, err.Error())
		return
	}
	h.successRequestResponse(w, http.StatusOK, team)
}
