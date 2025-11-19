package storage

import (
	"AvitoTest1/internal/models"
	"context"
	"errors"
	"fmt"

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
		//если пользователя не существует - returning вернет ошибку
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("client: user not found")
		}
		return nil, fmt.Errorf("server: %v", err)
	}
	return user, nil
}
