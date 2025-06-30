// Package postgres provides Postgres connection methods and common utils.
package postgres

import (
	"context"

	"github.com/jmoiron/sqlx"

	// Postgres driver.
	_ "github.com/lib/pq"
)

const driver = "postgres"

// Client is a wrapper to *sqlx.DB connection, extending with connection methods.
type Client struct {
	Conn *sqlx.DB
}

// New creates a Client instance
//
// Parameters:
//   - string: Postgres url (e.g.:"postgres://postgres:postgres@localhost:5432/pack?sslmode=disable")
//
// Returns :
//   - *Client: Postgres client wrapper instance
//   - error: if there is any error on connecting to database
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

// Ping do a database health check
//
// Parameters:
//   - context.Context: used to do the ping call, should be a context with timeout
//
// Returns:
//   - error: if there is error on pinging database
func (c *Client) Ping(ctx context.Context) error {
	return c.Conn.PingContext(ctx)
}

// Close closes the database connection with timeout
//
// Parameters:
//   - context.Context: used validate the action, should be a context with timeout
//
// Returns:
//   - error: if there is error on closing database connection
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
