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

	createNote := createnote.New(repo)
	archiveNote := archivenote.New(repo)

	r.Route("/api", func(api chi.Router) {
		api.Route("/v1", func(v1 chi.Router) {
			v1.Route("/notes", func(notes chi.Router) {
				notes.Post("/", createNote.HTTPHandler)
				notes.Get("/", listnotes.Route(repo))
				notes.Patch("/{id}", archiveNote.HTTPHandler)
			})
		})
	})

	return nil
}
