package postgres

import (
	"HATCH_APP/internal/note/domain"
	"HATCH_APP/pkg/store/postgres"

	"github.com/jmoiron/sqlx"
)

type TransactionManager struct {
	db *sqlx.DB
}

func NewTransactionManager(db *sqlx.DB) *TransactionManager {
	return &TransactionManager{db: db}
}

func (t *TransactionManager) Transact(fn func(input domain.TransactionManagerInput) error) error {
	return postgres.RunInTx(t.db, func(tx *sqlx.Tx) error {
		repo, err := NewNoteRepository(tx)
		if err != nil {
			return err
		}

		return fn(domain.TransactionManagerInput{
			NoteRepository: repo,
		})
	})
}
