package service

import (
	mye "AvitoTest1/internal/errors"
	"AvitoTest1/internal/models"
	"context"
	"log"
)

type UserServiceUserStorage interface {
	UpdateActive(ctx context.Context, userId string, status bool) (*models.User, error)
}
type UserServicePullRequestsStorage interface {
	SelectReviews(ctx context.Context, userId string) ([]*models.PullRequestShort, error)
}
type UserServiceImpl struct {
	Ust  UserServiceUserStorage
	PRst UserServicePullRequestsStorage
}

func NewUserService(usSt UserServiceUserStorage, prSt UserServicePullRequestsStorage) *UserServiceImpl {
	return &UserServiceImpl{
		Ust:  usSt,
		PRst: prSt,
	}
}
func (us *UserServiceImpl) SetIsActive(ctx context.Context, userid string, status bool) (*models.User, error) {
	user, err := us.Ust.UpdateActive(ctx, userid, status)
	if err != nil {
		log.Println(err)
		return nil, mye.ErrResourceNotFound
	}
	return user, nil
}
func (us *UserServiceImpl) GetUsersReview(ctx context.Context, userid string) ([]*models.PullRequestShort, error) {
	reqs, err := us.PRst.SelectReviews(ctx, userid)
	if err != nil {
		log.Println(err)
		return nil, mye.ErrResourceNotFound
	}
	return reqs, nil
}
