package service

import (
	"AvitoTest1/internal/models"
	"context"
	"errors"
	"log"
	"math/rand"
	"time"
)

type PullRequestServiceUserStorage interface {
	SelectActiveMembers(ctx context.Context, authorID string) (*models.Team, error)
}
type PullRequestServicePullRequestsStorage interface {
	InsertPullRequest(ctx context.Context, pr *models.PullRequest) error
}
type PullRequestServiceImpl struct {
	Ust  PullRequestServiceUserStorage
	PRst PullRequestServicePullRequestsStorage
}

func NewPullRequestService(ust PullRequestServiceUserStorage, prst PullRequestServicePullRequestsStorage) *PullRequestServiceImpl {
	return &PullRequestServiceImpl{
		Ust:  ust,
		PRst: prst,
	}
}
func (pr *PullRequestServiceImpl) PullRequestCreate(ctx context.Context, authorID string, id string, name string) (*models.PullRequest, error) {
	members, err := pr.Ust.SelectActiveMembers(ctx, authorID)
	if err != nil {
		log.Println(err)
		return nil, errors.New("resource not found")
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	count := r.Intn(3)
	var selectedMembers []models.TeamMember
	if count > 0 && len(members.Members) > 0 {
		shuffled := make([]models.TeamMember, len(members.Members))
		copy(shuffled, members.Members)

		rand.Shuffle(len(shuffled), func(i, j int) {
			shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
		})

		if count > len(shuffled) {
			count = len(shuffled)
		}
		selectedMembers = shuffled[:count]
	}
	var ids []string
	for _, mem := range selectedMembers {
		ids = append(ids, mem.UserId)
	}
	now := time.Now()
	prs := &models.PullRequest{
		AssignedReviewers: ids,
		AuthorId:          authorID,
		CreatedAt:         &now,
		PullRequestId:     id,
		PullRequestName:   name,
		Status:            models.PullRequestStatusOPEN,
	}
	err = pr.PRst.InsertPullRequest(ctx, prs)
	if err != nil {
		log.Println(err)
		return nil, errors.New("PR id already exists")
	}
	return prs, nil
}
