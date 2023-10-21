package usecase

import (
	"backend/internal/entity"
	"backend/internal/interfaces"
	"context"
)

type hashService interface {
	CheckHash(context.Context, []entity.Hash) ([]entity.Password, error)
	SavePassword(context.Context, entity.Password) error
}

type challengeService interface {
	GetChallenge(context.Context, int) (entity.Challenge, error)
}

type UseCase struct {
	log interfaces.Logger

	hashS hashService
	chgS  challengeService
}

func New(log interfaces.Logger, hashS hashService, chgS challengeService) *UseCase {
	return &UseCase{
		log:   log,
		hashS: hashS,
		chgS:  chgS,
	}
}
