package hashservice

import (
	"backend/internal/entity"
	"context"
	"fmt"
)

// CheckHash - проверяет, присутствуют ли данные хеши в базе данных.
func (hs *HashService) CheckHash(ctx context.Context, hashes []entity.Hash) ([]entity.Password, error) {
	var (
		err       error
		passwords = make([]entity.Password, 0)
		place     = "HashService CheckHash"
	)

	if passwords, err = hs.repo.CheckHash(ctx, hashes); err != nil {
		return nil, fmt.Errorf("%s: %w", place, err)
	}

	return passwords, nil
}
