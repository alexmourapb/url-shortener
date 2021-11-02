package stnpg

import (
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"

	_ "github.com/golang-migrate/migrate/v4/database/postgres" // Postgres migrate driver

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/source/httpfs"
)

// Migrations groups information about SQL migrations
type Migrations struct {
	// Folder name of the folder in which the migration .sql files are located
	Folder string

	// FS filesystem representing a migrations folder
	FS fs.FS
}

func RunMigrations(connString string, migrations Migrations) error {
	handler, err := GetMigrationHandler(connString, migrations)
	if err != nil {
		return err
	}

	handler.Log = &migrateLogger{logger: log.New(os.Stdout, "[migrations] ", log.LstdFlags)}

	if err := handler.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	srcErr, dbErr := handler.Close()
	if srcErr != nil {
		return fmt.Errorf("[migrations] failed to close DB source: %w", srcErr)
	}
	if dbErr != nil {
		return fmt.Errorf("[migrations] failed to close migrations repositories connection: %w", dbErr)
	}

	return nil
}

func GetMigrationHandler(connString string, migrations Migrations) (*migrate.Migrate, error) {
	if migrations.FS == nil {
		var err error
		handler, err := migrate.New("file://"+migrations.Folder, connString)
		if err != nil {
			return nil, fmt.Errorf("[migrations] failed to create migrate: %w", err)
		}

		return handler, nil
	}

	source, err := httpfs.New(http.FS(migrations.FS), migrations.Folder)
	if err != nil {
		return nil, fmt.Errorf("[migrations] failed to create httpfs driver: %w", err)
	}

	handler, err := migrate.NewWithSourceInstance("httpfs", source, connString)
	if err != nil {
		return nil, fmt.Errorf("[migrations] failed to create migrate source instance: %w", err)
	}

	return handler, nil
}

type migrateLogger struct {
	logger *log.Logger
}

func (l migrateLogger) Printf(arg string, vars ...interface{}) {
	l.logger.Printf(arg, vars...)
}

func (l migrateLogger) Verbose() bool {
	return true
}
