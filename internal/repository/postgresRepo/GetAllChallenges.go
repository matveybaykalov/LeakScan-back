package postgresrepo

import (
	"backend/internal/entity"
	"context"
	"fmt"
)

func (pgr *PostgresRepo) GetAllChallenges(ctx context.Context) ([]entity.Challenge, error) {
	var (
		err        error
		value      Challenges
		challenges = make([]entity.Challenge, 8)
		place      = "PostgresRepo GetAllChallenges"
	)

	if err = pgr.con.NewSelect().Model(&value).Limit(1).Scan(ctx); err != nil {
		return nil, fmt.Errorf("%s: %w", place, err)
	}

	if err = challenges[0].Set(value.Challenge0); err != nil {
		return nil, fmt.Errorf("%s: %w", place, err)
	}

	if err = challenges[1].Set(value.Challenge1); err != nil {
		return nil, fmt.Errorf("%s: %w", place, err)
	}

	if err = challenges[2].Set(value.Challenge2); err != nil {
		return nil, fmt.Errorf("%s: %w", place, err)
	}

	if err = challenges[3].Set(value.Challenge3); err != nil {
		return nil, fmt.Errorf("%s: %w", place, err)
	}

	if err = challenges[4].Set(value.Challenge4); err != nil {
		return nil, fmt.Errorf("%s: %w", place, err)
	}

	if err = challenges[5].Set(value.Challenge5); err != nil {
		return nil, fmt.Errorf("%s: %w", place, err)
	}

	if err = challenges[6].Set(value.Challenge6); err != nil {
		return nil, fmt.Errorf("%s: %w", place, err)
	}

	if err = challenges[7].Set(value.Challenge7); err != nil {
		return nil, fmt.Errorf("%s: %w", place, err)
	}

	return challenges, nil
}
