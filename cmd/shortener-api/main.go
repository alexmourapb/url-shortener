package main

import (
	"os"

	"github.com/rs/zerolog"

	"github.com/alexmourapb/url-shortener/app/common/configuration"
	mig "github.com/alexmourapb/url-shortener/app/common/storage/migrations"
	CommonPostgres "github.com/alexmourapb/url-shortener/app/common/storage/postgres"
	"github.com/alexmourapb/url-shortener/app/common/storage/redis"
	"github.com/alexmourapb/url-shortener/app/domain/shortener/usecase"
	"github.com/alexmourapb/url-shortener/app/gateways/db/postgres"
	"github.com/alexmourapb/url-shortener/app/gateways/db/postgres/shortener"
	shortener2 "github.com/alexmourapb/url-shortener/app/gateways/db/redis/shortener"
	"github.com/alexmourapb/url-shortener/app/gateways/http"
)

func main() {
	// Start Logger
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	// Load Config
	cfg, err := configuration.LoadConfig()
	if err != nil {
		logger.Fatal().Err(err).Msg("unable to load configuration")
	}
	log := logger.With().Str("app_name", cfg.Service.AppName).Logger()
	log = log.With().Str("environment", cfg.Service.Environment.String()).Logger()
	log.Info().Msg("logger successfully initialized")

	// Postgres
	pgConn, err := CommonPostgres.ConnectPool(cfg.Postgres.DSN(), CommonPostgres.LogLevelError)
	if err != nil {
		logger.Fatal().Err(err).Msg("unable to connect to postgres database")
	}
	defer pgConn.Close()
	log.Info().Msg("postgres connected successfully")

	// Migrations
	migration := mig.Migrations{
		Folder: "migrations",
		FS:     postgres.Migrations,
	}
	err = mig.RunMigrations(cfg.Postgres.URL(), migration)
	if err != nil {
		logger.Fatal().Err(err).Msg("unable to run migrations")
	}

	// New postgres repository
	repository := shortener.NewRepository(pgConn)

	//Redis
	redisPool, err := redis.InitPool(cfg.Redis)
	if err != nil {
		logger.Fatal().Err(err).Msg("error starting redis")
	}
	defer redisPool.Close()
	log.Info().Msg("redis connected successfully")

	// New cache repository
	cacheRepo := shortener2.NewCacheRepository(redisPool)

	// Shortener UseCase
	shortenerUseCase := usecase.NewShortenerUseCase(repository, cacheRepo)

	// Start api service
	api := http.NewShortenerApi(shortenerUseCase)
	api.Start(&log, cfg.Service.Host, cfg.Service.Port, cfg.Service.DestinationURLService)
}
