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
	repo, err := postgres.NewNoteRepository(db)
	if err != nil {
		return err
	}

	createNoteF := createnote.New(repo)
	archiveNoteF := archivenote.New(repo)
	listNotesF := listnotes.New(repo)

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
