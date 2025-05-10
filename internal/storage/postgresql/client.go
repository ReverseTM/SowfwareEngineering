package postgresql

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetConnPool(ctx context.Context, connString string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		return nil, err
	}

	return pool, nil
}
