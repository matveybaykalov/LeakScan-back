package clickhouserepo

import (
	"backend/internal/entity"
	"context"
	"fmt"
)

func (chr *ClickHouseRepo) GetAllChallenges(ctx context.Context) ([]entity.Challenge, error) {
	return nil, fmt.Errorf("implement ClickHouseRepo GetAllChallenges")
}
