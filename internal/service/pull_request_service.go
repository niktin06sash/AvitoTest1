package service

import (
	mye "AvitoTest1/internal/errors"
	"AvitoTest1/internal/models"
	"context"
	"log"
	"math/rand"
	"time"
)

type PullRequestServiceUserStorage interface {
	SelectActiveMembers(ctx context.Context, userID string) (*models.Team, error)
}
type PullRequestServicePullRequestsStorage interface {
	InsertPullRequest(ctx context.Context, pr *models.PullRequest) error
	UpdateStatusPullRequest(ctx context.Context, prID string, status models.PullRequestStatus) (*models.PullRequest, error)
	SelectPullRequest(ctx context.Context, prID string) (*models.PullRequest, error)
	UpdateReviewersPullRequest(ctx context.Context, pr_id string, rews []string) error
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
func (pr *PullRequestServiceImpl) CreatePullRequest(ctx context.Context, authorID string, id string, name string) (*models.PullRequest, error) {
	members, err := pr.Ust.SelectActiveMembers(ctx, authorID)
	if err != nil {
		log.Println(err)
		return nil, mye.ErrResourceNotFound
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
		return nil, mye.ErrPRExist
	}
	return prs, nil
}
func (pr *PullRequestServiceImpl) MergePullRequest(ctx context.Context, prID string) (*models.PullRequest, error) {
	prq, err := pr.PRst.UpdateStatusPullRequest(ctx, prID, models.PullRequestStatusMERGED)
	if err != nil {
		log.Println(err)
		return nil, mye.ErrResourceNotFound
	}
	return prq, nil
}
func (pr *PullRequestServiceImpl) ReassignPullRequest(ctx context.Context, olduserID string, prID string) (*models.PullRequest, string, error) {
	prq, err := pr.PRst.SelectPullRequest(ctx, prID)
	if err != nil {
		log.Println(err)
		return nil, "", mye.ErrResourceNotFound
	}
	if prq.Status == models.PullRequestStatusMERGED {
		return nil, "", mye.ErrMergedPR
	}
	delidx := -1
	for i, rew := range prq.AssignedReviewers {
		if rew == olduserID {
			delidx = i
			break
		}
	}
	if delidx == -1 {
		return nil, "", mye.ErrReviewerNotAssigned
	}
	members, err := pr.Ust.SelectActiveMembers(ctx, olduserID)
	if err != nil {
		log.Println(err)
		return nil, "", mye.ErrResourceNotFound
	}
	if len(members.Members) == 0 {
		return nil, "", mye.ErrNoActiveCandidate
	}
	//удаляем из доступных активных участников команды уже существующего ревьюера в пул-реквесте для избежания появления дубликата
	excluded := make(map[string]bool)
	for _, reviewerID := range prq.AssignedReviewers {
		excluded[reviewerID] = true
	}
	truecandidates := &models.Team{}
	for _, member := range members.Members {
		if !excluded[member.UserId] {
			truecandidates.Members = append(truecandidates.Members, member)
		}
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomidx := r.Intn(len(truecandidates.Members))
	randomreviewer := truecandidates.Members[randomidx]
	prq.AssignedReviewers[delidx] = randomreviewer.UserId
	err = pr.PRst.UpdateReviewersPullRequest(ctx, prID, prq.AssignedReviewers)
	if err != nil {
		log.Println(err)
		return nil, "", mye.ErrResourceNotFound
	}
	return prq, randomreviewer.UserId, nil
}
