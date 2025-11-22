package service

import (
	"AvitoTest1/internal/logger"
	"AvitoTest1/internal/service/pull_request"
	"AvitoTest1/internal/service/team"
	"AvitoTest1/internal/service/user"
)

type Service struct {
	UserService        *user.ServiceImpl
	TeamService        *team.ServiceImpl
	PullRequestService *pull_request.ServiceImpl
}

func NewService(logger *logger.Logger, prus pull_request.UserStorage, prpr pull_request.PullRequestsStorage, tts team.TeamStorage, tsu team.UserStorage, txman team.TxManagerStorage, ususs user.UserStorage, usprs user.PullRequestsStorage) *Service {
	return &Service{
		UserService:        user.NewUserService(logger, ususs, usprs),
		TeamService:        team.NewTeamService(logger, tts, tsu, txman),
		PullRequestService: pull_request.NewPullRequestService(logger, prus, prpr),
	}
}
