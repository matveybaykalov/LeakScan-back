package clickhouserepo

import (
	"backend/internal/entity"
	"context"
	"fmt"
)

func (chr *ClickHouseRepo) GetChallenge(ctx context.Context, n int) (entity.Challenge, error) {
	return entity.Challenge{}, fmt.Errorf("implement ClickHouseRepo GetChallenge")
}
