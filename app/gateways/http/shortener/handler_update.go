package shortener

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/alexmourapb/url-shortener/app/common/api/responses"
	"github.com/alexmourapb/url-shortener/app/common/instrumentation"
	"github.com/alexmourapb/url-shortener/app/common/shared"
	"github.com/alexmourapb/url-shortener/app/domain/shortener/vos"
	"github.com/alexmourapb/url-shortener/app/gateways/http/shortener/models"
	_ "github.com/alexmourapb/url-shortener/docs/swagger"
)

// HandlerUpdate ...
// @Tags Update
// @Summary Update a short url
// @Description Update a short url
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param Body body models.UpdateRequest true "Body"
// @Success 200
// @Failure 400 {object} swagger.ErrRespBadRequest
// @Failure 404 {object} swagger.ErrRespNotFound
// @Failure 500 {object} swagger.ErrRespInternalServerError
// @Router /api/v1/shortener/{id} [put]
func (h Handler) HandlerUpdate(w http.ResponseWriter, r *http.Request) {
	operation := "Handler.Get"
	log, err := instrumentation.LogFromContext(r.Context(), *h.logger, operation)
	if err != nil {
		log.Err(err)
		_ = responses.Send(w, responses.ErrInternalServerError, http.StatusInternalServerError)
		return
	}
	log.Info().Msg("starting update short url")

	vars := mux.Vars(r)
	id := vars["id"]

	var shortenerRequest models.UpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&shortenerRequest); err != nil {
		title := "empty body"
		log.Err(err).Msg(title)
		responseError := responses.FullError{
			Type:  responses.ErrBadRequest.Type,
			Title: title,
		}
		_ = responses.Send(w, responseError, http.StatusBadRequest)
		return
	}

	if err := h.Validator.Validate(shortenerRequest); err != nil {
		title := "invalid request body"
		log.Err(err).Msg(title)
		responseError := responses.FullError{
			Type:  responses.ErrBadRequest.Type,
			Title: err.Error(),
		}
		_ = responses.Send(w, responseError, http.StatusBadRequest)
		return
	}

	input := vos.UpdateInput{
		ID:     id,
		URL:    shortenerRequest.URL,
		Active: shortenerRequest.Active,
	}

	err = h.UseCase.Update(r.Context(), log, input)
	if err != nil {
		title := "failed to update short url"
		log.Err(err).Msg(title)
		switch {
		case errors.Is(err, shared.ErrURLNotFound):
			_ = responses.Send(w, responses.ErrNotFound, http.StatusNotFound)
		default:
			_ = responses.Send(w, responses.ErrInternalServerError, http.StatusInternalServerError)
		}
		return
	}

	log.Info().Msg("update short url successfully")
	_ = responses.Send(w, nil, http.StatusOK)
}
