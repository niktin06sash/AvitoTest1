package models

import "time"

//перенес модели из package handler, чтобы использовать в сервисе

// PullRequestStatus defines model for PullRequest.Status.
type PullRequestStatus string

// PullRequestShort defines model for PullRequestShort.
// PullRequestShortStatus defines model for PullRequestShort.Status.
type PullRequestShortStatus string

// Defines values for PullRequestStatus.
const (
	PullRequestStatusMERGED PullRequestStatus = "MERGED"
	PullRequestStatusOPEN   PullRequestStatus = "OPEN"
)

// Defines values for PullRequestShortStatus.
const (
	PullRequestShortStatusMERGED PullRequestShortStatus = "MERGED"
	PullRequestShortStatusOPEN   PullRequestShortStatus = "OPEN"
)

type PullRequest struct {
	// AssignedReviewers user_id назначенных ревьюверов (0..2)
	AssignedReviewers []string          `json:"assigned_reviewers"`
	AuthorId          string            `json:"author_id"`
	CreatedAt         *time.Time        `json:"createdAt"`
	MergedAt          *time.Time        `json:"mergedAt"`
	PullRequestId     string            `json:"pull_request_id"`
	PullRequestName   string            `json:"pull_request_name"`
	Status            PullRequestStatus `json:"status"`
}
type PullRequestShort struct {
	AuthorId        string                 `json:"author_id"`
	PullRequestId   string                 `json:"pull_request_id"`
	PullRequestName string                 `json:"pull_request_name"`
	Status          PullRequestShortStatus `json:"status"`
}

// Team defines model for Team.
type Team struct {
	Members  []TeamMember `json:"members"`
	TeamName string       `json:"team_name"`
}

// TeamMember defines model for TeamMember.
type TeamMember struct {
	IsActive bool   `json:"is_active"`
	UserId   string `json:"user_id"`
	Username string `json:"username"`
}

// User defines model for User.
type User struct {
	IsActive bool   `json:"is_active"`
	TeamName string `json:"team_name"`
	UserId   string `json:"user_id"`
	Username string `json:"username"`
}
