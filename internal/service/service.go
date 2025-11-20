package service

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Service struct {
	UserService        *UserServiceImpl
	TeamService        *TeamServiceImpl
	PullRequestService *PullRequestServiceImpl
}

func NewService(txman TxManagerStorage, tts TeamServiceTeamStorage, tsu TeamServiceUserStorage, ususs UserServiceUserStorage, usprs UserServicePullRequestsStorage, prus PullRequestServiceUserStorage, prpr PullRequestServicePullRequestsStorage) *Service {
	return &Service{
		UserService:        NewUserService(ususs, usprs),
		TeamService:        NewTeamService(tts, tsu, txman),
		PullRequestService: NewPullRequestService(prus, prpr),
	}
}

type TxManagerStorage interface {
	BeginTx(ctx context.Context) (pgx.Tx, error)
	Commit(ctx context.Context, tx pgx.Tx) error
	Rollback(ctx context.Context, tx pgx.Tx) error
}
