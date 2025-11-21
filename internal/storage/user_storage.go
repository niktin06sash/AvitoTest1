package storage

import (
	"AvitoTest1/internal/models"
	"context"

	"github.com/jackc/pgx/v5"
)

type UserStorageImpl struct {
	db *DBObject
}

func NewUserStorage(db *DBObject) *UserStorageImpl {
	return &UserStorageImpl{db: db}
}
func (us *UserStorageImpl) UpdateActive(ctx context.Context, userId string, status bool) (*models.User, error) {
	const query = `
	UPDATE users
	SET is_active = $1
	WHERE user_id = $2
	RETURNING user_id, username, team_name, is_active;
	`
	user := &models.User{}
	err := us.db.pool.QueryRow(ctx, query, status, userId).Scan(
		&user.UserId, &user.Username, &user.TeamName, &user.IsActive)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (ts *UserStorageImpl) InsertOrUpdateUsers(ctx context.Context, tx pgx.Tx, members []models.TeamMember, teamname string) error {
	const query = `
    INSERT INTO users (user_id, username, is_active, team_name)
    VALUES ($1, $2, $3, $4)
    ON CONFLICT (user_id) 
    DO UPDATE SET 
        username = EXCLUDED.username,
        is_active = EXCLUDED.is_active,
        team_name = EXCLUDED.team_name
    `
	for _, member := range members {
		_, err := tx.Exec(ctx, query,
			member.UserId,
			member.Username,
			member.IsActive,
			teamname,
		)
		if err != nil {
			return err
		}
	}
	return nil
}
func (us *UserStorageImpl) SelectTeamMember(ctx context.Context, teamname string) (*models.Team, error) {
	const query = `
    SELECT user_id, username, is_active
    FROM users 
    WHERE team_name = $1
    `
	rows, err := us.db.pool.Query(ctx, query, teamname)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	team := &models.Team{}
	for rows.Next() {
		member := models.TeamMember{}
		err := rows.Scan(&member.UserId, &member.Username, &member.IsActive)
		if err != nil {
			return nil, err
		}
		team.Members = append(team.Members, member)
	}
	team.TeamName = teamname
	return team, nil
}
func (us *UserStorageImpl) SelectActiveMembers(ctx context.Context, userID string) (*models.Team, error) {
	const query = `
    SELECT u2.user_id, u2.username, u2.is_active, u1.team_name
    FROM users AS u1
    JOIN users AS u2 
        ON u2.team_name = u1.team_name
       AND u2.is_active = TRUE
       AND u2.user_id <> u1.user_id
    WHERE u1.user_id = $1
    `
	rows, err := us.db.pool.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	team := &models.Team{}
	for rows.Next() {
		member := models.TeamMember{}
		err := rows.Scan(&member.UserId, &member.Username, &member.IsActive, &team.TeamName)
		if err != nil {
			return nil, err
		}
		team.Members = append(team.Members, member)
	}
	return team, nil
}
