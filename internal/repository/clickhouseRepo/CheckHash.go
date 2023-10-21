package clickhouserepo

import (
	"backend/internal/entity"
	"context"
	"fmt"
)

func (chr *ClickHouseRepo) CheckHash(ctx context.Context, hashes []entity.Hash) ([]entity.Password, error) {
	return nil, fmt.Errorf("implement ClickHouseRepo CheckHash")
}
