package challengeservice

import (
	"backend/internal/entity"
	"context"
	"fmt"
)

func (cs *ChallengeService) GetChallenge(ctx context.Context, n int) (entity.Challenge, error) {
	var (
		err       error
		challenge entity.Challenge
		place     = "ChallengeService GetChallenge"
	)

	if challenge, err = cs.repo.GetChallenge(ctx, n); err != nil {
		return entity.Challenge{}, fmt.Errorf("%s: %w", place, err)
	}

	return challenge, nil
}
