package service

import (
	"AvitoTest1/internal/models"
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

type TeamServiceTeamStorage interface {
	InsertTeam(ctx context.Context, tx pgx.Tx, team *models.Team) error
	SelectExistTeam(ctx context.Context, tn string) error
}
type TeamServiceUserStorage interface {
	InsertOrUpdateUsers(ctx context.Context, tx pgx.Tx, members []models.TeamMember, teamname string) error
	SelectTeamMember(ctx context.Context, tn string) (*models.Team, error)
}
type TeamServiceImpl struct {
	Tst TeamServiceTeamStorage
	Txm TxManagerStorage
	Ust TeamServiceUserStorage
}

func NewTeamService(tst TeamServiceTeamStorage, ust TeamServiceUserStorage, txman TxManagerStorage) *TeamServiceImpl {
	return &TeamServiceImpl{
		Tst: tst,
		Txm: txman,
		Ust: ust,
	}
}
func (ts *TeamServiceImpl) AddTeam(ctx context.Context, team *models.Team) error {
	tx, err := ts.Txm.BeginTx(ctx)
	if err != nil {
		log.Println(err)
		return err
	}
	err = ts.Tst.InsertTeam(ctx, tx, team)
	if err != nil {
		log.Println(err)
		rolbackerr := ts.Txm.Rollback(ctx, tx)
		if rolbackerr != nil {
			log.Println(rolbackerr)
		}
		return err
	}
	err = ts.Ust.InsertOrUpdateUsers(ctx, tx, team.Members, team.TeamName)
	if err != nil {
		log.Println(err)
		rolbackerr := ts.Txm.Rollback(ctx, tx)
		if rolbackerr != nil {
			log.Println(rolbackerr)
		}
		return err
	}
	err = ts.Txm.Commit(ctx, tx)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
func (ts *TeamServiceImpl) GetTeam(ctx context.Context, teamname string) (*models.Team, error) {
	err := ts.Tst.SelectExistTeam(ctx, teamname)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	members, err := ts.Ust.SelectTeamMember(ctx, teamname)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return members, nil
}
