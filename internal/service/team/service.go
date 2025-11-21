package team

import (
	mye "AvitoTest1/internal/errors"
	"AvitoTest1/internal/models"
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

type TeamStorage interface {
	InsertTeam(ctx context.Context, tx pgx.Tx, team *models.Team) error
	SelectExistTeam(ctx context.Context, tn string) error
}
type UserStorage interface {
	InsertOrUpdateUsers(ctx context.Context, tx pgx.Tx, members []models.TeamMember, teamname string) error
	SelectTeamMember(ctx context.Context, tn string) (*models.Team, error)
}
type TxManagerStorage interface {
	BeginTx(ctx context.Context) (pgx.Tx, error)
	Commit(ctx context.Context, tx pgx.Tx) error
	Rollback(ctx context.Context, tx pgx.Tx) error
}

type ServiceImpl struct {
	Tst TeamStorage
	Txm TxManagerStorage
	Ust UserStorage
}

func NewTeamService(tst TeamStorage, ust UserStorage, txman TxManagerStorage) *ServiceImpl {
	return &ServiceImpl{
		Tst: tst,
		Txm: txman,
		Ust: ust,
	}
}
func (ts *ServiceImpl) AddTeam(ctx context.Context, team *models.Team) error {
	tx, err := ts.Txm.BeginTx(ctx)
	if err != nil {
		log.Println(err)
		return mye.ErrTeamExist
	}
	err = ts.Tst.InsertTeam(ctx, tx, team)
	if err != nil {
		log.Println(err)
		rolbackerr := ts.Txm.Rollback(ctx, tx)
		if rolbackerr != nil {
			log.Println(rolbackerr)
		}
		return mye.ErrTeamExist
	}
	err = ts.Ust.InsertOrUpdateUsers(ctx, tx, team.Members, team.TeamName)
	if err != nil {
		log.Println(err)
		rolbackerr := ts.Txm.Rollback(ctx, tx)
		if rolbackerr != nil {
			log.Println(rolbackerr)
		}
		return mye.ErrTeamExist
	}
	err = ts.Txm.Commit(ctx, tx)
	if err != nil {
		log.Println(err)
		return mye.ErrTeamExist
	}
	return nil
}
func (ts *ServiceImpl) GetTeam(ctx context.Context, teamname string) (*models.Team, error) {
	err := ts.Tst.SelectExistTeam(ctx, teamname)
	if err != nil {
		log.Println(err)
		return nil, mye.ErrTeamExist
	}
	members, err := ts.Ust.SelectTeamMember(ctx, teamname)
	if err != nil {
		log.Println(err)
		return nil, mye.ErrTeamExist
	}
	return members, nil
}
