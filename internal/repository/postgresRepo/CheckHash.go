package postgresrepo

import (
	"backend/internal/entity"
	"context"
	"fmt"
)

func (pgr *PostgresRepo) CheckHash(ctx context.Context, hashes []entity.Hash) ([]entity.Password, error) {
	var (
		err    error
		values = []string{}
		passes = []entity.Password{}
		place  = "PostgresRepo CheckHash"
	)

	if len(hashes) > 8 || len(hashes) < 1 {
		err = fmt.Errorf("hash count must be between 1 and 8")
		return nil, fmt.Errorf("%s: %w", place, err)
	}

	sql := pgr.con.NewSelect().ModelTableExpr("hash").Column("password")

	for i, hash := range hashes {
		sql.Where("chg_? = ?", i, hash.Get())
	}

	if err = sql.Scan(context.Background(), &values); err != nil {
		return nil, fmt.Errorf("%s: %w", place, err)
	}

	for _, value := range values {
		pass := entity.Password{}

		if err = pass.Set(value); err != nil {
			return nil, fmt.Errorf("%s: %w", place, err)
		}

		passes = append(passes, pass)
	}

	return passes, nil
}
