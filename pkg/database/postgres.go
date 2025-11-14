package database

import (
	"context"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

const postgresgDriver = "postgres"

var ErrQueryPreparation = errors.New("query preparation error")

func NewPostgresClient(ctx context.Context, url string) (*sqlx.DB, error) {
	ctx, stop := context.WithTimeout(ctx, 10*time.Second)
	defer stop()

	db, err := sqlx.ConnectContext(ctx, postgresgDriver, url)
	if err != nil {
		return nil, err
	}

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
