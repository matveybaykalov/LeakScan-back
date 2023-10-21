package usecase

import (
	"backend/internal/entity"
	"context"
	"fmt"
)

func (uc *UseCase) Start(ctx context.Context) (entity.Challenge, error) {
	var (
		err       error
		challenge entity.Challenge
		place     = "UseCase Start"
	)

	if challenge, err = uc.chgS.GetChallenge(ctx, 0); err != nil {
		return entity.Challenge{}, fmt.Errorf("%s: %w", place, err)
	}

	return challenge, nil
}
