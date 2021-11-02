package postgres

import (
	"embed"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
)

//go:embed migrations
var Migrations embed.FS
