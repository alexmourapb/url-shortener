package shortener

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/alexmourapb/url-shortener/app/common/api/responses"
	"github.com/alexmourapb/url-shortener/app/common/instrumentation"
	"github.com/alexmourapb/url-shortener/app/gateways/http/shortener/models"
	_ "github.com/alexmourapb/url-shortener/docs/swagger"
)

// HandlerGet ...
// @Summary Get a short url
// @Description Get a short url
// @Tags Get a short url
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 201 {object} models.GetResponse
// @Failure 400 {object} swagger.ErrRespBadRequest
// @Failure 500 {object} swagger.ErrRespInternalServerError
// @Router /api/v1/shortener/{id} [get]
func (h Handler) HandlerGet(w http.ResponseWriter, r *http.Request) {
	operation := "Handler.Get"
	log, err := instrumentation.LogFromContext(r.Context(), *h.logger, operation)
	log.Info().Msg("starting get short url")

	vars := mux.Vars(r)
	id := vars["id"]

	output, err := h.UseCase.Get(r.Context(), id)
	if err != nil {
		title := "failed to get short url"
		log.Err(err).Msg(title)
		_ = responses.Send(w, responses.ErrInternalServerError, http.StatusInternalServerError)
		return
	}

	response := &models.GetResponse{
		URL: output,
	}

	log.Info().Msg("get short url successfully")
	_ = responses.Send(w, response, http.StatusCreated)
}
