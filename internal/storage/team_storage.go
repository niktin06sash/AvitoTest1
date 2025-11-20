package storage

import (
	"AvitoTest1/internal/models"
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

type TeamStorageImpl struct {
	db *DBObject
}

func NewTeamStorage(db *DBObject) *TeamStorageImpl {
	return &TeamStorageImpl{db: db}
}
func (ts *TeamStorageImpl) InsertTeam(ctx context.Context, tx pgx.Tx, team *models.Team) error {
	const query = `
	INSERT INTO teams (team_name) VALUES ($1) ON CONFLICT (team_name) DO NOTHING RETURNING team_name
	`
	var insertedName string
	err := tx.QueryRow(ctx, query, team.TeamName).Scan(&insertedName)
	return err
}
func (ts *TeamStorageImpl) SelectExistTeam(ctx context.Context, tn string) error {
	var exist bool
	const query = `
	SELECT EXISTS(SELECT 1 FROM teams WHERE team_name = $1)
	`
	err := ts.db.pool.QueryRow(ctx, query, tn).Scan(&exist)
	if err != nil {
		return err
	}
	if !exist {
		return errors.New("command not found")
	}
	return nil
}
