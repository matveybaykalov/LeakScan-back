package postgresrepo

import (
	"backend/internal/entity"
	"context"
	"fmt"
)

func (pgr *PostgresRepo) GetChallenge(ctx context.Context, n int) (entity.Challenge, error) {
	var (
		err       error
		challenge entity.Challenge
		value     int
		place     = "PostgresRepo GetChallenge"
	)

	if n > 8 {
		err = fmt.Errorf("n must be less than or equal to 7")
		return entity.Challenge{}, fmt.Errorf("%s: %w", place, err)
	}

	if err = pgr.con.NewSelect().ModelTableExpr("challenge").
		ColumnExpr("chg_?", n).Limit(1).Scan(ctx, &value); err != nil {
		return entity.Challenge{}, fmt.Errorf("%s: %w", place, err)
	}

	challenge.Set(value)

	return challenge, nil
}
