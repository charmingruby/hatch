package postgres

import (
	"context"
	"database/sql"
	"time"

	"HATCH_APP/examples/packagebyfeat/core"
	"HATCH_APP/pkg/db/postgres"

	"github.com/jmoiron/sqlx"
)

type NoteRepository struct {
	db    *sqlx.DB
	stmts map[string]*sqlx.Stmt
}

func NewNoteRepository(db *sqlx.DB) (*NoteRepository, error) {
	stmts := make(map[string]*sqlx.Stmt)

	for queryName, statement := range noteQueries() {
		stmt, err := db.Preparex(statement)
		if err != nil {
			return nil,
				postgres.NewPreparationErr(queryName, "note", err)
		}

		stmts[queryName] = stmt
	}

	return &NoteRepository{
		db:    db,
		stmts: stmts,
	}, nil
}

func (r *NoteRepository) statement(queryName string) (*sqlx.Stmt, error) {
	stmt, ok := r.stmts[queryName]

	if !ok {
		return nil,
			postgres.NewStatementNotPreparedErr(queryName, "note")
	}

	return stmt, nil
}

func (r *NoteRepository) Create(ctx context.Context, note core.Note) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	stmt, err := r.statement(createNote)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx,
		&note.ID,
		&note.Title,
		&note.Content,
		&note.Archived,
		&note.CreatedAt,
		&note.UpdatedAt,
	)

	return err
}

func (r *NoteRepository) FindByID(ctx context.Context, id string) (core.Note, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	stmt, err := r.statement(findNoteByID)
	if err != nil {
		return core.Note{}, err
	}

	var note core.Note

	if err := stmt.QueryRowContext(ctx, id).Scan(
		&note.ID,
		&note.Title,
		&note.Content,
		&note.Archived,
		&note.CreatedAt,
		&note.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return core.Note{}, nil
		}

		return core.Note{}, err
	}

	return note, nil
}

func (r *NoteRepository) List(ctx context.Context) ([]core.Note, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	stmt, err := r.statement(listNotes)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.QueryxContext(ctx)
	if err != nil {
		return nil, err
	}

	var notes []core.Note

	for rows.Next() {
		var note core.Note
		if err := rows.StructScan(&note); err != nil {
			return nil, err
		}

		notes = append(notes, note)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return notes, nil
}

func (r *NoteRepository) Save(ctx context.Context, note core.Note) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	stmt, err := r.statement(saveNote)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx,
		&note.Archived,
		&note.UpdatedAt,
		&note.ID,
	)

	return err
}
