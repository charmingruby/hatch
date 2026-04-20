package container

import (
	"HATCH_APP/db/migration"
	"HATCH_APP/pkg/connection/postgres"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang-migrate/migrate/v4"
	pgMg "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jmoiron/sqlx"
	pg "github.com/testcontainers/testcontainers-go/modules/postgres"
)

func SetupPostgres(t *testing.T) (*sqlx.DB, func()) {
	ctx := context.Background()

	container, err := pg.Run(ctx,
		"postgres:15-alpine",
		pg.WithDatabase("testdb"),
		pg.WithUsername("user"),
		pg.WithPassword("pass"),
	)
	if err != nil {
		t.Fatalf("failed to start container: %v", err)
	}

	connStr, err := container.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		t.Fatalf("failed to get connection string: %v", err)
	}

	var db *sqlx.DB
	for range 10 {
		db, err = postgres.Connect(ctx, connStr)
		if err == nil {
			break
		}

		time.Sleep(500 * time.Millisecond)
	}
	if err != nil {
		if err := container.Terminate(ctx); err != nil {
			t.Logf("failed to terminate postgres: %v", err)
		}

		t.Fatalf("failed to connect to postgres: %v", err)
	}

	runMigrations(t, db)

	teardown := func() {
		if err := db.Close(); err != nil {
			t.Logf("failed to close postgres connection: %v", err)
		}

		if err := container.Terminate(ctx); err != nil {
			t.Logf("failed to terminate postgres: %v", err)
		}
	}

	return db, teardown
}

func runMigrations(t *testing.T, db *sqlx.DB) {
	driver, err := pgMg.WithInstance(db.DB, &pgMg.Config{})
	if err != nil {
		t.Fatalf("failed to create migration driver: %v", err)
	}

	sourceDriver, err := iofs.New(migration.Files, ".")
	if err != nil {
		t.Fatalf("failed to init iofs source: %v", err)
	}

	m, err := migrate.NewWithInstance(
		"iofs",
		sourceDriver,
		"postgres",
		driver,
	)
	if err != nil {
		t.Fatalf("failed to init migrate: %v", err)
	}

	if _, _, err := m.Version(); err == nil {
		if err := m.Down(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			t.Fatalf("failed to run down migrations: %v", err)
		}
	}

	if err := m.Up(); err != nil {
		t.Fatalf("failed to run migrations: %v", err)
	}
}
