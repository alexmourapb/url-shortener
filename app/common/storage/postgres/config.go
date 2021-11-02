package postgres

import "fmt"

type Config struct {
	DatabaseName string `envconfig:"DATABASE_NAME" default:"developer"`
	User         string `envconfig:"DATABASE_USER" default:"postgres"`
	Password     string `envconfig:"DATABASE_PASSWORD" default:"postgres"`
	Host         string `envconfig:"DATABASE_HOST_DIRECT" default:"localhost"`
	Port         string `envconfig:"DATABASE_PORT_DIRECT" default:"5432"`
	PoolMinSize  string `envconfig:"DATABASE_POOL_MIN_SIZE" default:"2"`
	PoolMaxSize  string `envconfig:"DATABASE_POOL_MAX_SIZE" default:"10"`
	SSLMode      string `envconfig:"DATABASE_SSLMODE" default:"disable"`
	SSLRootCert  string `envconfig:"DATABASE_SSL_ROOTCERT"`
	SSLCert      string `envconfig:"DATABASE_SSL_CERT"`
	SSLKey       string `envconfig:"DATABASE_SSL_KEY"`
}

func (c Config) DSN() string {
	connectString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s pool_min_conns=%s pool_max_conns=%s",
		c.User, c.Password, c.Host, c.Port, c.DatabaseName, c.PoolMinSize, c.PoolMaxSize)

	if c.SSLMode != "" {
		connectString = fmt.Sprintf("%s sslmode=%s",
			connectString, c.SSLMode)
	}

	if c.SSLRootCert != "" {
		connectString = fmt.Sprintf("%s sslrootcert=%s sslcert=%s sslkey=%s",
			connectString, c.SSLRootCert, c.SSLCert, c.SSLKey)
	}

	return connectString
}

func (c Config) URL() string {
	if c.SSLMode == "" {
		c.SSLMode = "disable"
	}

	connectString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.User, c.Password, c.Host, c.Port, c.DatabaseName, c.SSLMode)

	if c.SSLRootCert != "" {
		connectString = fmt.Sprintf("%s&sslrootcert=%s&sslcert=%s&sslkey=%s",
			connectString, c.SSLRootCert, c.SSLCert, c.SSLKey)
	}

	return connectString
}
