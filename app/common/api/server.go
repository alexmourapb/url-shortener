package api

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog"
)

// ListenAndServe with Graceful Shutdown
func ListenAndServe(endpoint string, logger *zerolog.Logger, handler http.Handler, writeTimeout time.Duration, readTimeout time.Duration) {
	log := logger.With().Str("SERVER", "Listen and Serve").Logger()
	srv := &http.Server{
		Handler:      handler,
		Addr:         endpoint,
		WriteTimeout: writeTimeout,
		ReadTimeout:  readTimeout,
	}

	c := make(chan os.Signal, 1)
	idleConnections := make(chan struct{})
	signal.Notify(c, os.Interrupt, syscall.SIGINT)
	go func() {
		<-c
		// create context with timeout
		ctx, cancel := context.WithTimeout(context.Background(), writeTimeout)
		defer cancel()

		// start http shutdown
		if err := srv.Shutdown(ctx); err != nil {
			log.Error().AnErr("shutdown", err)
		}

		close(idleConnections)
	}()

	log.Info().Msg("listening at " + endpoint)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal().Msg("listen and Serve fail " + err.Error())
	}

	log.Info().Msg("waiting idle connections...")
	<-idleConnections
	log.Info().Msg("bye bye")
}
