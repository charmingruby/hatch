package postgres

import (
	"context"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

const driver = "postgres"

type Client struct {
	Conn *sqlx.DB
}

func New(ctx context.Context, url string) (*Client, error) {
	db, err := sqlx.ConnectContext(ctx, driver, url)
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
