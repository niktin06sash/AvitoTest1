package service

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Service struct {
	UserService *UserServiceImpl
	TeamService *TeamServiceImpl
}

func NewService(usSt UserStorage, txman TxManagerStorage, tts TeamServiceTeamStorage, tsu TeamServiceUserStorage) *Service {
	return &Service{
		UserService: NewUserService(usSt),
		TeamService: NewTeamService(tts, tsu, txman),
	}
}

type TxManagerStorage interface {
	BeginTx(ctx context.Context) (pgx.Tx, error)
	Commit(ctx context.Context, tx pgx.Tx) error
	Rollback(ctx context.Context, tx pgx.Tx) error
}
