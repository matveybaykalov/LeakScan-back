package postgresrepo

import (
	"backend/internal/entity"
	"context"
	"fmt"
)

func (pgr *PostgresRepo) SaveHash(ctx context.Context, pass entity.Password, hashes []entity.Hash) error {
	var (
		err   error
		value Hash
		place = "PostgresRepo SaveHash"
	)

	if len(hashes) != 8 {
		return fmt.Errorf("hashes count must be equal to 7")
	}

	value.Password = pass.Get()

	value.Challenge0 = hashes[0].Get()
	value.Challenge1 = hashes[1].Get()
	value.Challenge2 = hashes[2].Get()
	value.Challenge3 = hashes[3].Get()
	value.Challenge4 = hashes[4].Get()
	value.Challenge5 = hashes[5].Get()
	value.Challenge6 = hashes[6].Get()
	value.Challenge7 = hashes[7].Get()

	if _, err = pgr.con.NewInsert().Model(&value).Exec(ctx); err != nil {
		return fmt.Errorf("%s: %w", place, err)
	}

	return nil
}
