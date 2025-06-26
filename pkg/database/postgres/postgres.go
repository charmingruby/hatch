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

func New(url string) (*Client, error) {
	db, err := sqlx.Connect(driver, url)
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

func (c *Client) Close(ctx context.Context) error {
	if err := c.Conn.Close(); err != nil {
		return err
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		return nil
	}
}
