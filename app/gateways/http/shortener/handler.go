package shortener

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"

	"github.com/alexmourapb/url-shortener/app/common/validator"
	"github.com/alexmourapb/url-shortener/app/domain/shortener"
)

type Handler struct {
	UseCase   shortener.UseCase
	Validator *validator.JSONValidator
	logger    *zerolog.Logger
}

func NewHandler(public *mux.Router, logger *zerolog.Logger, useCase shortener.UseCase) *Handler {
	h := &Handler{
		UseCase:   useCase,
		Validator: validator.NewJSONValidator(),
		logger:    logger,
	}

	// Public routes
	publicShortener := public.PathPrefix("/shortener").Subrouter()

	// Create Short URL
	publicShortener.HandleFunc("/", h.HandlerCreate).Methods(http.MethodPost)

	// Get URL
	publicShortener.HandleFunc("/{id}", h.HandlerGet).Methods(http.MethodGet)

	return h
}
