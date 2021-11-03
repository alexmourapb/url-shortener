package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/urfave/negroni"

	"github.com/alexmourapb/url-shortener/app/common/api"
	"github.com/alexmourapb/url-shortener/app/domain/shortener/usecase"
	"github.com/alexmourapb/url-shortener/app/gateways/http/healthcheck"
	"github.com/alexmourapb/url-shortener/app/gateways/http/middleware"
	"github.com/alexmourapb/url-shortener/app/gateways/http/shortener"
	_ "github.com/alexmourapb/url-shortener/app/gateways/http/shortener/models"
	_ "github.com/alexmourapb/url-shortener/docs/swagger"
)

var (
	writeTimeout = 15 * time.Second
	readTimeout  = 15 * time.Second
)

type ShortenerApi struct {
	UseCase     *usecase.Shortener
	HealthCheck *healthcheck.Handler
}

// NewShortenerApi ...
// @title Shortener URL Service
// @version 1.0
// @description Documentation for Shortener-URL Api
func NewShortenerApi(useCase *usecase.Shortener) *ShortenerApi {
	return &ShortenerApi{
		UseCase:     useCase,
		HealthCheck: healthcheck.NewHandler(),
	}
}

func (c *ShortenerApi) Start(logger *zerolog.Logger, host, port, srvDest string) {
	// Router
	r := mux.NewRouter()

	// Handlers
	r.PathPrefix("/docs/v1/swagger/").Handler(httpSwagger.WrapHandler).Methods(http.MethodGet)
	r.HandleFunc("/healthcheck", c.HealthCheck.Get).Methods(http.MethodGet)

	publicV1 := r.PathPrefix("/api/v1").Subrouter()

	// Shortener Routes
	shortener.NewHandler(publicV1, logger, c.UseCase, srvDest)

	// Negroni to manage middlewares
	recovery := negroni.NewRecovery()
	recovery.PrintStack = false
	n := negroni.New(recovery)
	n.UseFunc(middleware.AddHeaders)
	n.UseHandler(middleware.Cors(r))

	endpoint := fmt.Sprintf("%s:%s", host, port)
	api.ListenAndServe(endpoint, logger, n, writeTimeout, readTimeout)
}
