package shortener

import (
	pgx "github.com/jackc/pgx/v4/pgxpool"

	"github.com/alexmourapb/url-shortener/app/domain/shortener"
)

var _ shortener.Repository = &Repository{}

type Repository struct {
	*pgx.Pool
}

func NewRepository(db *pgx.Pool) *Repository {
	return &Repository{
		Pool: db,
	}
}
