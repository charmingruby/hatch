package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/charmingruby/pack/internal/note/model"
	"github.com/charmingruby/pack/pkg/database/postgres"
	"github.com/jmoiron/sqlx"
)

type NoteRepo struct {
	db    *sqlx.DB
	stmts map[string]*sqlx.Stmt
}

func NewNoteRepo(db *sqlx.DB) (*NoteRepo, error) {
	stmts := make(map[string]*sqlx.Stmt)

	for queryName, statement := range noteQueries() {
		stmt, err := db.Preparex(statement)
		if err != nil {
			return nil,
				postgres.NewPreparationErr(queryName, "note", err)
		}

		stmts[queryName] = stmt
	}

	return &NoteRepo{
		db:    db,
		stmts: stmts,
	}, nil
}

func (r *NoteRepo) statement(queryName string) (*sqlx.Stmt, error) {
	stmt, ok := r.stmts[queryName]

	if !ok {
		return nil,
			postgres.NewStatementNotPreparedErr(queryName, "note")
	}

	return stmt, nil
}

func (r *NoteRepo) Create(ctx context.Context, note model.Note) error {
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

func (r *NoteRepo) FindByID(ctx context.Context, id string) (model.Note, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	stmt, err := r.statement(findNoteByID)
	if err != nil {
		return model.Note{}, err
	}

	var note model.Note

	if err := stmt.QueryRowContext(ctx, id).Scan(
		&note.ID,
		&note.Title,
		&note.Content,
		&note.Archived,
		&note.CreatedAt,
		&note.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return model.Note{}, nil
		}

		return model.Note{}, err
	}

	return note, nil
}

func (r *NoteRepo) List(ctx context.Context) ([]model.Note, error) {
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

	var notes []model.Note

	for rows.Next() {
		var note model.Note
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

func (r *NoteRepo) Save(ctx context.Context, note model.Note) error {
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
