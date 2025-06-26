package postgres

import (
	"context"
	"github/charmingruby/pack/internal/device/model"
	"github/charmingruby/pack/pkg/database/postgres"
	"time"

	"github.com/jmoiron/sqlx"
)

const (
	defaultQueryTimeout = 10

	createDevice = "create device"
)

type DeviceRepo struct {
	db    *sqlx.DB
	stmts map[string]*sqlx.Stmt
}

func NewDeviceRepo(db *sqlx.DB) (*DeviceRepo, error) {
	stmts := make(map[string]*sqlx.Stmt)

	for queryName, statement := range deviceQueries() {
		stmt, err := db.Preparex(statement)
		if err != nil {
			return nil, postgres.NewPreparationErr(queryName, "blocked source", err)
		}

		stmts[queryName] = stmt
	}

	return &DeviceRepo{
		db:    db,
		stmts: stmts,
	}, nil
}

func deviceQueries() map[string]string {
	return map[string]string{
		createDevice: `
		INSERT INTO devices(
			id, hardware_id, hardware_type, created_at
		) VALUES(
			$1, $2, $3, $4
		)`,
	}
}

func (r *DeviceRepo) statement(queryName string) (*sqlx.Stmt, error) {
	stmt, ok := r.stmts[queryName]

	if !ok {
		return nil,
			postgres.NewStatementNotPreparedErr(queryName, "blocked source")
	}

	return stmt, nil
}

func (r *DeviceRepo) Create(ctx context.Context, device model.Device) error {
	stmt, err := r.statement(createDevice)
	if err != nil {
		return postgres.NewStatementNotPreparedErr(createDevice, "device")
	}

	ctx, cancel := context.WithTimeout(ctx, defaultQueryTimeout*time.Second)
	defer cancel()

	if _, err := stmt.ExecContext(
		ctx,
		device.ID,
		device.HardwareID,
		device.HardwareType,
		device.CreatedAt,
	); err != nil {
		return err
	}

	return nil
}
