package hashservice

import (
	"backend/internal/entity"
	"backend/internal/interfaces"
	"context"
)

type Repo interface {
	SaveHash(context.Context, entity.Password, []entity.Hash) error
	CheckHash(context.Context, []entity.Hash) ([]entity.Password, error)
	GetAllChallenges(context.Context) ([]entity.Challenge, error)
}

type HashService struct {
	repo Repo
	log  interfaces.Logger
}

func New(log interfaces.Logger, repo Repo) *HashService {
	return &HashService{
		repo: repo,
		log:  log,
	}
}
