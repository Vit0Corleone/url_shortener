package redirect

import (
	"errors"
	"log/slog"
	"net/http"
	resp "url_short/internal/lib/api/response"
	"url_short/internal/lib/logger/sl"
	"url_short/internal/storage"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type URLGetter interface {
	GetUrl(alias string) (string, error)
}

func New(log *slog.Logger, getUrl URLGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.redirect.New"

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

		resURL, err := getUrl.GetUrl(alias)
		if err != nil {
			if errors.Is(err, storage.ErrUrlNotFound) {
				log.Info("url not found", "alias", alias)

				render.JSON(w, r, resp.Error("not found"))

				return
			}

			log.Error("failed to get url", sl.Err(err))

			render.JSON(w, r, resp.Error("internal error"))

			return
		}

		log.Info("got url", slog.String("url", resURL))

		// redirect to found url
		http.Redirect(w, r, resURL, http.StatusFound)
	}
}
