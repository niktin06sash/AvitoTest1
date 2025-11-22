package user

import (
	mye "AvitoTest1/internal/errors"
	"AvitoTest1/internal/logger"
	"AvitoTest1/internal/models"
	"context"

	"go.uber.org/zap"
)

type UserStorage interface {
	UpdateActive(ctx context.Context, userId string, status bool) (*models.User, error)
}
type PullRequestsStorage interface {
	SelectReviews(ctx context.Context, userId string) ([]*models.PullRequestShort, error)
}
type ServiceImpl struct {
	Ust    UserStorage
	PRst   PullRequestsStorage
	Logger *logger.Logger
}

func NewUserService(logger *logger.Logger, usSt UserStorage, prSt PullRequestsStorage) *ServiceImpl {
	return &ServiceImpl{
		Ust:  usSt,
		PRst: prSt,
	}
}
func (us *ServiceImpl) SetIsActive(ctx context.Context, userid string, status bool) (*models.User, error) {
	user, err := us.Ust.UpdateActive(ctx, userid, status)
	if err != nil {
		us.Logger.ZapLogger.Error("UpdateActive", zap.Error(err), zap.String("user_id", userid))
		return nil, mye.ErrResourceNotFound
	}
	us.Logger.ZapLogger.Info("Successful set active", zap.String("user_id", userid), zap.Bool("status", status))
	return user, nil
}
func (us *ServiceImpl) GetUsersReview(ctx context.Context, userid string) ([]*models.PullRequestShort, error) {
	reqs, err := us.PRst.SelectReviews(ctx, userid)
	if err != nil {
		us.Logger.ZapLogger.Error("SelectReviews", zap.Error(err), zap.String("user_id", userid))
		return nil, mye.ErrResourceNotFound
	}
	us.Logger.ZapLogger.Info("Successful get user's reviews", zap.String("user_id", userid))
	return reqs, nil
}
