package note

import (
	"HATCH_APP/internal/note/feature/archivenote"
	"HATCH_APP/internal/note/feature/createnote"
	"HATCH_APP/internal/note/feature/listnotes"
	"HATCH_APP/internal/note/infra/store/postgres"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func Register(r chi.Router, db *sqlx.DB) error {
	noteRepo, err := postgres.NewNoteRepository(db)
	if err != nil {
		return err
	}

	createNoteF := createnote.New(noteRepo)
	archiveNoteF := archivenote.New(noteRepo)
	listNotesF := listnotes.New(noteRepo)

	r.Route("/v1/notes", func(r chi.Router) {
		r.Post("/", createNoteF.CreateNoteEndpoint)
		r.Get("/", listNotesF.ListNotesEndpoint)
		r.Patch("/{id}", archiveNoteF.ArchiveNoteEndpoint)
	})

	return nil
}
