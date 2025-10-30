package http

import (
	"HATCH_APP/examples/packagebyfeat/core"
	"HATCH_APP/internal/shared/errs"
	"HATCH_APP/internal/shared/http/rest"
	"HATCH_APP/pkg/telemetry"
	"errors"

	"github.com/gin-gonic/gin"
)

type CreateNoteRequest struct {
	Title   string `json:"title"   binding:"required" validate:"required,gt=0"`
	Content string `json:"content" binding:"required" validate:"required,gt=0"`
}

func CreateNoteHandler(log *telemetry.Logger, uc core.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		log.InfoContext(ctx, "endpoint/CreateNote: request received")

		req, err := rest.ParseRequest[CreateNoteRequest](c)
		if err != nil {
			log.ErrorContext(
				ctx,
				"endpoint/CreateNote: unable to parse payload",
				"error", err,
			)

			rest.SendBadRequestResponse(c, err.Error())
		}

		id, err := uc.CreateNote(ctx, req.Title, req.Content)
		if err != nil {
			var databaseErr *errs.DatabaseError
			if errors.As(err, &databaseErr) {
				log.ErrorContext(
					ctx,
					"endpoint/CreateNote: database error",
					"error", databaseErr.Unwrap(),
				)

				rest.SendInternalServerErrorResponse(c)
				return
			}

			log.ErrorContext(
				ctx,
				"endpoint/CreateNote: unknown error", "error", err,
			)

			rest.SendInternalServerErrorResponse(c)
			return
		}

		log.InfoContext(ctx, "endpoint/CreateNote: finished successfully")

		rest.SendCreatedResponse(c, id, "note")
	}
}

type FetchNotesResponse = []core.Note

func FetchNotesHandler(log *telemetry.Logger, uc core.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		log.InfoContext(ctx, "endpoint/FetchNotes: request received")

		notes, err := uc.FetchNotes(ctx)
		if err != nil {
			var databaseErr *errs.DatabaseError
			if errors.As(err, &databaseErr) {
				log.ErrorContext(
					ctx,
					"endpoint/FetchNotes: database error",
					"error", databaseErr.Unwrap(),
				)

				rest.SendInternalServerErrorResponse(c)
				return
			}

			log.ErrorContext(
				ctx,
				"endpoint/FetchNotes: unknown error", "error", err,
			)

			rest.SendInternalServerErrorResponse(c)
			return
		}

		var res = notes

		log.InfoContext(ctx, "endpoint/FetchNotes: finished successfully")

		rest.SendOKResponse(
			c,
			"",
			res,
		)
	}
}

func ArchiveNoteHandler(log *telemetry.Logger, uc core.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		log.InfoContext(ctx, "endpoint/ArchiveNote: request received")

		id := c.Param("id")

		if err := uc.ArchiveNote(ctx, id); err != nil {
			var notFoundErr *errs.NotFoundError
			if errors.As(err, &notFoundErr) {
				log.ErrorContext(
					ctx,
					"endpoint/ArchiveNote: not found error",
					"error", err,
				)

				rest.SendNotFoundResponse(c, err.Error())
				return
			}

			var databaseErr *errs.DatabaseError
			if errors.As(err, &databaseErr) {
				log.ErrorContext(
					ctx,
					"endpoint/ArchiveNote: database error",
					"error", databaseErr.Unwrap().Error(),
				)

				rest.SendInternalServerErrorResponse(c)
				return
			}

			log.ErrorContext(
				ctx,
				"endpoint/ArchiveNote: unknown error", "error", err,
			)

			rest.SendInternalServerErrorResponse(c)
			return
		}

		log.InfoContext(ctx, "endpoint/ArchiveNote: finished successfully")

		rest.SendEmptyResponse(c)
	}
}
