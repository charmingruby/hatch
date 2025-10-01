package postgres

import (
	"context"

	"HATCH_APP/config"

	"github.com/jmoiron/sqlx"
	"go.uber.org/fx"

	_ "github.com/lib/pq"
)

const driver = "postgres"

type Client struct {
	Conn *sqlx.DB
}

func New(cfg *config.Config) (*Client, error) {
	db, err := sqlx.Connect(driver, cfg.PostgresURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Client{Conn: db}, nil
}

func (c *Client) Ping(ctx context.Context) error {
	return c.Conn.PingContext(ctx)
}

func (c *Client) Close() error {
	if err := c.Conn.Close(); err != nil {
		return err
	}

	return nil
}

var Module = fx.Module("postgres",
	fx.Provide(New),
	fx.Provide(func(c *Client) *sqlx.DB {
		return c.Conn
	}),
	fx.Invoke(func(lc fx.Lifecycle, db *Client) {
		lc.Append(fx.Hook{
			OnStop: func(ctx context.Context) error {
				return db.Close()
			},
		})
	}),
)
