package postgres

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

const driver = "postgres"

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
