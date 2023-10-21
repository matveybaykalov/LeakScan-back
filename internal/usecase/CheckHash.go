package usecase

import (
	"backend/internal/entity"
	"context"
	"fmt"
)

// Возвращаемое bool значение показывает, стоит ли клиенту делать ещё один запрос
// Возможны несколько сценариев:
//	1) возникла какая-то ошибка при работе с БД
//	2) не нашлось ни одного пароля -> продолжать проверки не нужно
//	3) нашлось несколько паролей их число не болше некоторого максимального ->
//		продолжать исследование не нужно, вернуть найденные пароли пользователю
//	4) нашлось слишком много паролей -> требуются дополнительные проверки
func (uc *UseCase) CheckHash(ctx context.Context, hashes []entity.Hash) (bool, entity.Challenge, []entity.Password, error) {
	var (
		err          error
		passwords    = make([]entity.Password, 0)
		challenge    entity.Challenge
		maxPasswords = 5
		place        = "UseCase CheckHash"
	)

	if passwords, err = uc.hashS.CheckHash(ctx, hashes); err != nil {
		return false, entity.Challenge{}, nil, fmt.Errorf("%s: %w", place, err)
	}

	if len(passwords) == 0 {
		return false, entity.Challenge{}, nil, nil
	}

	if len(passwords) <= maxPasswords {
		return false, entity.Challenge{}, passwords, nil
	}

	nextChallenge := len(hashes)
	if challenge, err = uc.chgS.GetChallenge(ctx, nextChallenge); err != nil {
		return false, entity.Challenge{}, nil, fmt.Errorf("%s: %w", place, err)
	}

	return true, challenge, nil, nil
}
