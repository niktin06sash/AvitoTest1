package user

import (
	mye "AvitoTest1/internal/errors"
	"AvitoTest1/internal/models"
	"context"
	"log"
)

type UserStorage interface {
	UpdateActive(ctx context.Context, userId string, status bool) (*models.User, error)
}
type PullRequestsStorage interface {
	SelectReviews(ctx context.Context, userId string) ([]*models.PullRequestShort, error)
}
type ServiceImpl struct {
	Ust  UserStorage
	PRst PullRequestsStorage
}

func NewUserService(usSt UserStorage, prSt PullRequestsStorage) *ServiceImpl {
	return &ServiceImpl{
		Ust:  usSt,
		PRst: prSt,
	}
}
func (us *ServiceImpl) SetIsActive(ctx context.Context, userid string, status bool) (*models.User, error) {
	user, err := us.Ust.UpdateActive(ctx, userid, status)
	if err != nil {
		log.Println(err)
		return nil, mye.ErrResourceNotFound
	}
	return user, nil
}
func (us *ServiceImpl) GetUsersReview(ctx context.Context, userid string) ([]*models.PullRequestShort, error) {
	reqs, err := us.PRst.SelectReviews(ctx, userid)
	if err != nil {
		log.Println(err)
		return nil, mye.ErrResourceNotFound
	}
	return reqs, nil
}
