package storage

import (
	"AvitoTest1/internal/models"
	"context"
)

type TeamStorageImpl struct {
	db *DBObject
}

func NewTeamStorage(db *DBObject) *TeamStorageImpl {
	return &TeamStorageImpl{db: db}
}
func (ts *TeamStorageImpl) InsertTeam(ctx context.Context, team *models.Team) error {
	const query = `
	INSERT INTO teams (team_name) VALUES ($1) ON CONFLICT (team_name) DO NOTHING RETURNING team_name
	`
	var insertedName string
	err := ts.db.pool.QueryRow(ctx, query, team.TeamName).Scan(&insertedName)
	return err
}
