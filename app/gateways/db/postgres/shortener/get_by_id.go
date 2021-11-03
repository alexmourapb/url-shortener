package shortener

import (
	"context"

	pgxType "github.com/jackc/pgx/v4"

	"github.com/alexmourapb/url-shortener/app/common/shared"
	"github.com/alexmourapb/url-shortener/app/domain/shortener/entities"
)

func (r *Repository) GetByID(ctx context.Context, id string) (*entities.ShortURL, error) {
	var shortURL entities.ShortURL

	query := "SELECT id, url, active FROM short_urls WHERE id = $1"

	err := r.Pool.QueryRow(ctx, query, id).Scan(
		&shortURL.ID,
		&shortURL.URL,
		&shortURL.Active,
	)
	if err != nil {
		if err == pgxType.ErrNoRows {
			return nil, shared.ErrURLNotFound
		}
		return nil, err
	}

	return &shortURL, nil
}
