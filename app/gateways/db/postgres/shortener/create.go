package shortener

import (
	"context"
)

func (r *Repository) Save(ctx context.Context, id, url string) error {
	query := `INSERT INTO
		short_urls (
			id,
			url
		) VALUES($1, $2)`

	_, err := r.Pool.Exec(ctx, query, id, url)
	if err != nil {
		return err
	}

	return nil
}
