package postgres

import (
	"context"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

const driver = "postgres"

var ErrQueryPreparation = errors.New("query preparation error")

type Querier interface {
	sqlx.ExtContext
	Preparex(query string) (*sqlx.Stmt, error)
}

func Connect(ctx context.Context, url string) (*sqlx.DB, error) {
	ctx, stop := context.WithTimeout(ctx, 10*time.Second)
	defer stop()

	db, err := sqlx.ConnectContext(ctx, driver, url)
	if err != nil {
		return nil, err
	}

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
