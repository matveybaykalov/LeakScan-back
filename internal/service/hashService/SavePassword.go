package hashservice

import (
	"backend/internal/entity"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
)

func (hs *HashService) SavePassword(ctx context.Context, password entity.Password) error {
	var (
		err        error
		challenges = make([]entity.Challenge, 0)
		hashes     = make([]entity.Hash, 0)
		place      = "HashService SavePassword"
		hasher     hash.Hash
	)

	if challenges, err = hs.repo.GetAllChallenges(ctx); err != nil {
		return fmt.Errorf("%s: %w", place, err)
	}

	hasher = sha256.New()
	hasher.Write([]byte(password.Get()))

	previous := hex.EncodeToString(hasher.Sum(nil))
	length := len(previous) / 4
	previous = previous[0:length]

	for _, challenge := range challenges {
		hasher = sha256.New()

		input := fmt.Sprintf("%s%d", previous, challenge.Get())
		hasher.Write([]byte(input))

		newHash := entity.Hash{}
		if err = newHash.Set(hex.EncodeToString(hasher.Sum(nil))[0:length]); err != nil {
			return fmt.Errorf("%s: %w", place, err)
		}
		hashes = append(hashes, newHash)

		previous = newHash.Get()
	}

	if err = hs.repo.SaveHash(ctx, password, hashes); err != nil {
		return err
	}

	return nil
}
