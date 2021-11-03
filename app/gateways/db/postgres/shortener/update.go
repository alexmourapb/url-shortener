package shortener

import (
	"context"
	"time"

	"github.com/alexmourapb/url-shortener/app/domain/shortener/entities"
)

func (r *Repository) Update(ctx context.Context, input entities.ShortURL) error {
	query := `UPDATE
			short_urls SET
				url = $1,
				active = $2,
				updated_at = $3
			WHERE id = $4`

	_, err := r.Exec(ctx, query,
		input.URL,
		input.Active,
		time.Now().UTC(),
		input.ID,
	)
	if err != nil {
		return err
	}

	return nil
}
