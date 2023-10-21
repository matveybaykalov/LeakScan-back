package clickhouserepo

import (
	"backend/internal/entity"
	"context"
	"fmt"
)

func (chr *ClickHouseRepo) SaveHash(ctx context.Context, hashes []entity.Hash) error {
	return fmt.Errorf("implement ClickHouseRepo SaveHash")
}
