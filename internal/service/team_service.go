package service

import (
	"AvitoTest1/internal/models"
	"context"
	"log"
)

type TeamServiceTeamStorage interface {
	InsertTeam(ctx context.Context, team *models.Team) error
}
type TeamServiceUserStorage interface {
	InsertOrUpdateUsers(ctx context.Context, members []models.TeamMember, teamname string) error
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
	//добавить транзакцию
	err := ts.Tst.InsertTeam(ctx, team)
	if err != nil {
		log.Println(err)
		return err
	}
	err = ts.Ust.InsertOrUpdateUsers(ctx, team.Members, team.TeamName)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
