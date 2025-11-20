package storage

import (
	"AvitoTest1/internal/models"
	"context"
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
func (ts *UserStorageImpl) InsertOrUpdateUsers(ctx context.Context, members []models.TeamMember, teamname string) error {
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
		_, err := ts.db.pool.Exec(ctx, query,
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
func (us *UserStorageImpl) SelectReviews(ctx context.Context, userId string) ([]*models.PullRequestShort, error) {
	const query = `
	SELECT pull_request_id, pull_request_name, author_id, status
	FROM pull_requests
	WHERE $1 = ANY(pr.assigned_reviewers);
	`
	rows, err := us.db.pool.Query(ctx, query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	prs := make([]*models.PullRequestShort, 0)
	for rows.Next() {
		pr := &models.PullRequestShort{}
		err := rows.Scan(&pr.PullRequestId, &pr.PullRequestName, &pr.AuthorId, &pr.Status)
		if err != nil {
			return nil, err
		}
		prs = append(prs, pr)
	}
	return prs, nil
}
