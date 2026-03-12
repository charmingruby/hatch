package httpx

import (
	"HATCH_APP/pkg/o11y"
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

type External struct {
	DB *sqlx.DB
}

func registerProbes(r *chi.Mux, ext External) {
	r.Get("/api/livez", livenessRoute())
	r.Get("/api/readyz", readinessRoute(ext))
}

func livenessRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		WriteOKResponse(w, "healthy", nil)
	}
}

func readinessRoute(ext External) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log := o11y.FromContext(r.Context())

		ctx := r.Context()

		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		if err := ext.DB.PingContext(ctx); err != nil {
			log.ErrorContext(ctx, "endpoint/Readiness: database error", "error", err)

			WriteServiceUnavailableResponse(w, "database")
			return
		}

		WriteOKResponse(w, "ready", nil)
	}
}
