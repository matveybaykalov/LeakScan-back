package usecase

import (
	"backend/internal/entity"
	"context"
	"fmt"
)

func (uc *UseCase) AddPassword(ctx context.Context, password entity.Password) error {
	var (
		err   error
		place = "UseCase AddPassword"
	)

	if err = uc.hashS.SavePassword(ctx, password); err != nil {
		return fmt.Errorf("%s: %w", place, err)
	}

	return nil
}
