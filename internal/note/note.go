package note

import (
	"HATCH_APP/internal/note/feature/archivenote"
	"HATCH_APP/internal/note/feature/createnote"
	"HATCH_APP/internal/note/feature/listnotes"
	"HATCH_APP/internal/note/infra/database/postgres"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func Register(r *chi.Mux, db *sqlx.DB) error {
	noteRepo, err := postgres.NewNoteRepository(db)
	if err != nil {
		return err
	}

	createNoteF := createnote.New(noteRepo)
	archiveNoteF := archivenote.New(noteRepo)
	listNotesF := listnotes.New(noteRepo)

	r.Route("/api", func(apiR chi.Router) {
		apiR.Route("/v1", func(v1R chi.Router) {
			v1R.Route("/notes", func(notesR chi.Router) {
				notesR.Post("/", createNoteF.CreateNoteEndpoint)
				notesR.Get("/", listNotesF.ListNotesEndpoint)
				notesR.Patch("/{id}", archiveNoteF.ArchiveNoteEndpoint)
			})
		})
	})

	return nil
}
