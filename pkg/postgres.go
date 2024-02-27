package pkg

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Pg_Client interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)

	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row

	Begin(ctx context.Context) (pgx.Tx, error)
}

func NewPostgresClient() (pool *pgxpool.Pool, err error) {
	ctx := context.Background()
	dsn1 := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", "postgres", "postgres", "localhost", "5435", "platadb")
	// dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Dbname)

	pool, err = pgxpool.New(ctx, dsn1)
	if err != nil {
		panic("sdsdsdsds")
		return nil, err
	}

	return pool, err
}
