package shortener

import (
	"encoding/json"
	"net/http"

	"github.com/alexmourapb/url-shortener/app/common/api/responses"
	"github.com/alexmourapb/url-shortener/app/common/instrumentation"
	"github.com/alexmourapb/url-shortener/app/gateways/http/shortener/models"
	_ "github.com/alexmourapb/url-shortener/docs/swagger"
)

// HandlerCreate ...
// @Summary Create a new short url
// @Description Create a new short url
// @Tags Create a new short url
// @Accept json
// @Produce json
// @Param Body body models.CreateRequest true "Body"
// @Success 201 {object} models.CreateResponse
// @Failure 400 {object} swagger.ErrRespBadRequest
// @Failure 500 {object} swagger.ErrRespInternalServerError
// @Router /api/v1/shortener [post]
func (h Handler) HandlerCreate(w http.ResponseWriter, r *http.Request) {
	operation := "Handler.Create"
	log, err := instrumentation.LogFromContext(r.Context(), *h.logger, operation)
	log.Info().Msg("starting create short url")

	var shortenerRequest models.CreateRequest
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

	output, err := h.UseCase.Create(r.Context(), shortenerRequest.URL)
	if err != nil {
		title := "failed to create short url"
		log.Err(err).Msg(title)
		_ = responses.Send(w, responses.ErrInternalServerError, http.StatusInternalServerError)
		return
	}

	response := &models.CreateResponse{
		URL: output,
	}

	log.Info().Msg("short url created successfully")
	_ = responses.Send(w, response, http.StatusCreated)
}
