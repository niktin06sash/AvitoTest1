package service

import (
	"AvitoTest1/internal/service/pull_request"
	"AvitoTest1/internal/service/team"
	"AvitoTest1/internal/service/user"
)

type Service struct {
	UserService        *user.ServiceImpl
	TeamService        *team.ServiceImpl
	PullRequestService *pull_request.ServiceImpl
}

func NewService(prus pull_request.UserStorage, prpr pull_request.PullRequestsStorage, tts team.TeamStorage, tsu team.UserStorage, txman team.TxManagerStorage, ususs user.UserStorage, usprs user.PullRequestsStorage) *Service {
	return &Service{
		UserService:        user.NewUserService(ususs, usprs),
		TeamService:        team.NewTeamService(tts, tsu, txman),
		PullRequestService: pull_request.NewPullRequestService(prus, prpr),
	}
}
