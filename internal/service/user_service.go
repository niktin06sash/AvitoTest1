package service

import (
	"AvitoTest1/internal/models"
	"context"
	"log"
)

type UserStorage interface {
	UpdateActive(ctx context.Context, userId string, status bool) (*models.User, error)
}
type UserServiceImpl struct {
	Ust UserStorage
}

func NewUserService(usSt UserStorage) *UserServiceImpl {
	return &UserServiceImpl{
		Ust: usSt,
	}
}
func (us *UserServiceImpl) SetIsActive(ctx context.Context, userid string, status bool) (*models.User, error) {
	user, err := us.Ust.UpdateActive(ctx, userid, status)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return user, nil
}
