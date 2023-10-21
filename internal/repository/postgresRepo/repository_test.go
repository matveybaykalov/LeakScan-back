package postgresrepo

import (
	"backend/internal/entity"
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type test struct {
	name   string
	input  int
	output int
}

func getRepo(t *testing.T) *PostgresRepo {
	var pgDSN = os.Getenv("PG_DSN")

	require.NotEmpty(t, pgDSN, "Please, set PG_DSN env var for testing")

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(pgDSN)))
	pgConnection := bun.NewDB(sqldb, pgdialect.New())

	pgRepo := New(pgConnection)
	return pgRepo
}

func TestGetChallenge(t *testing.T) {
	pgRepo := getRepo(t)

	tests := []test{
		{
			name:   "test getChallenge #0",
			input:  0,
			output: 760812383,
		},
		{
			name:   "test getChallenge #1",
			input:  1,
			output: 674013276,
		},
		{
			name:   "test getChallenge #2",
			input:  2,
			output: 196833153,
		},
		{
			name:   "test getChallenge #3",
			input:  3,
			output: 542033844,
		},
		{
			name:   "test getChallenge #4",
			input:  4,
			output: 391622057,
		},
		{
			name:   "test getChallenge #5",
			input:  5,
			output: 285878725,
		},
		{
			name:   "test getChallenge #6",
			input:  6,
			output: 173960679,
		},
		{
			name:   "test getChallenge #7",
			input:  7,
			output: 136639778,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			challenge, err := pgRepo.GetChallenge(context.Background(), tc.input)
			require.NoError(t, err)

			require.Equal(t, tc.output, challenge.Get())
		})
	}
}

func TestGetAllChallenges(t *testing.T) {
	pgRepo := getRepo(t)

	challenges, err := pgRepo.GetAllChallenges(context.Background())
	require.NoError(t, err)

	fmt.Println(challenges)
}

func TestSaveHash(t *testing.T) {
	pgRepo := getRepo(t)

	pass := entity.Password{}
	err := pass.Set("olol")
	require.NoError(t, err)

	hashes := []entity.Hash{}

	for i := 0; i < 8; i += 1 {
		hs := entity.Hash{}

		err = hs.Set(fmt.Sprintf("%s %d", "hash", i))
		require.NoError(t, err)

		hashes = append(hashes, hs)
	}

	err = pgRepo.SaveHash(context.Background(), pass, hashes)
	require.NoError(t, err)
}

func TestCheckHash(t *testing.T) {
	pgRepo := getRepo(t)

	hashes := []entity.Hash{}

	h1 := entity.Hash{}
	err := h1.Set("hash 0")
	require.NoError(t, err)
	hashes = append(hashes, h1)

	passes, err := pgRepo.CheckHash(context.Background(), hashes)
	require.NoError(t, err)

	fmt.Println(passes)
}
