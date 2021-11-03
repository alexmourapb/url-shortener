package shortener

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/alexmourapb/url-shortener/app/common/api/responses"
	"github.com/alexmourapb/url-shortener/app/common/instrumentation"
	"github.com/alexmourapb/url-shortener/app/common/shared"
	_ "github.com/alexmourapb/url-shortener/docs/swagger"
)

// HandlerGet ...
// @Summary Redirect by a short url
// @Description Redirect by a short url
// @Tags Redirect by a short url
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 308
// @Failure 404 {object} swagger.ErrRespNotFound
// @Failure 500 {object} swagger.ErrRespInternalServerError
// @Router /api/v1/shortener/{id} [get]
func (h Handler) HandlerGet(w http.ResponseWriter, r *http.Request) {
	operation := "Handler.Get"
	log, err := instrumentation.LogFromContext(r.Context(), *h.logger, operation)
	log.Info().Msg("starting get short url")

	vars := mux.Vars(r)
	id := vars["id"]

	url, err := h.UseCase.Get(r.Context(), log, id)
	if err != nil {
		title := "failed to get short url"
		log.Err(err).Msg(title)
		switch {
		case errors.Is(err, shared.ErrURLNotFound):
			_ = responses.Send(w, responses.ErrNotFound, http.StatusNotFound)
		default:
			_ = responses.Send(w, responses.ErrInternalServerError, http.StatusInternalServerError)
		}
		return
	}

	log.Info().Msg("get short url successfully")
	http.Redirect(w, r, url, http.StatusPermanentRedirect)
}
