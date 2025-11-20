package storage

import (
	"AvitoTest1/internal/models"
	"context"
)

type PullRequestStorageImpl struct {
	db *DBObject
}

func NewPullRequestStorage(db *DBObject) *PullRequestStorageImpl {
	return &PullRequestStorageImpl{db: db}
}
func (us *PullRequestStorageImpl) SelectReviews(ctx context.Context, userId string) ([]*models.PullRequestShort, error) {
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
func (us *PullRequestStorageImpl) InsertPullRequest(ctx context.Context, pr *models.PullRequest) error {
	const query = `
		INSERT INTO pull_requests (assigned_reviewers, author_id, createdAt, pull_request_id, pull_request_name, status)
	`
	_, err := us.db.pool.Exec(ctx, query, pr.AssignedReviewers, pr.AuthorId, pr.CreatedAt, pr.PullRequestId, pr.PullRequestName, pr.Status)
	return err
}
func (us *PullRequestStorageImpl) UpdateStatusPullRequest(ctx context.Context, prID string, status models.PullRequestStatus) (*models.PullRequest, error) {
	const query = `UPDATE pull_requests 
	SET 
   		status = $2,
    	merged_at = COALESCE(merged_at, NOW())
	WHERE pr_id = $1 
	RETURNING status, merged_at, assigned_reviewers, author_id, status, pull_request_id, pull_request_name, created_at
	`
	pr := &models.PullRequest{}
	err := us.db.pool.QueryRow(ctx, query, prID, status).Scan(&pr.Status, &pr.MergedAt, &pr.AssignedReviewers, &pr.AuthorId, &pr.Status,
		&pr.PullRequestId, &pr.PullRequestName, &pr.CreatedAt)
	if err != nil {
		return nil, err
	}
	return pr, nil
}
