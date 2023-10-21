package challengeservice

import (
	"backend/internal/entity"
	"backend/internal/interfaces"
	"context"
)

type Repo interface {
	GetChallenge(context.Context, int) (entity.Challenge, error)
}

type ChallengeService struct {
	repo Repo
	log  interfaces.Logger
}

func New(log interfaces.Logger, repo Repo) *ChallengeService {
	return &ChallengeService{
		repo: repo,
		log:  log,
	}
}
