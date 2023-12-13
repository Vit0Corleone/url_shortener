package delete

import (
	"log/slog"
	"net/http"
	resp "url_short/internal/lib/api/response"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type Deleter interface {
	Delete(alias string) error
}

func New(log *slog.Logger, deleteAlias Deleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.delete.New"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		alias := chi.URLParam(r, "alias")
		if alias == "" {
			log.Info("alias is empty")

			render.JSON(w, r, resp.Error("invalid request"))

			return
		}

		if err := deleteAlias.Delete(alias); err != nil {
			log.Info("could not delete an alias")

			render.JSON(w, r, resp.Error("nothing to delete"))

			return
		}

		log.Info("successfully deleted", slog.String("alias", alias))

		render.JSON(w, r, "deleted")
	}
}
