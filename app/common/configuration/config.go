package configuration

import (
	"github.com/kelseyhightower/envconfig"

	"github.com/alexmourapb/url-shortener/app/common/environment"
	commonPostgres "github.com/alexmourapb/url-shortener/app/common/storage/postgres"
	"github.com/alexmourapb/url-shortener/app/common/storage/redis"
)

// Config defines the service configuration
type Config struct {
	Service ServiceConfig

	//Storage
	Postgres commonPostgres.Config
	Redis    redis.Config
}

type ServiceConfig struct {
	AppName     string                  `envconfig:"APP_NAME" default:"shortener-api"`
	Host        string                  `envconfig:"APP_HOST" default:"0.0.0.0"`
	Port        string                  `envconfig:"API_PORT" default:"8000"`
	Environment environment.Environment `envconfig:"ENVIRONMENT" default:"developer"`
	SwaggerHost string                  `envconfig:"SWAGGER_HOST" default:"http://localhost:8000"`
}

func LoadConfig() (*Config, error) {
	var config Config
	noPrefix := ""

	err := envconfig.Process(noPrefix, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
