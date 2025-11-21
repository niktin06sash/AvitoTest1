package errors

import "errors"

var (
	ErrResourceNotFound    = errors.New("resource not found")
	ErrPRExist             = errors.New("PR id already exists")
	ErrReviewerNotAssigned = errors.New("reviewer is not assigned to this PR")
	ErrNoActiveCandidate   = errors.New("no active replacement candidate in team")
	ErrMergedPR            = errors.New("cannot reassign on merged PR")
	ErrTeamExist           = errors.New("team_name already exists")
)
