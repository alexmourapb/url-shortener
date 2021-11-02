package shortener

import (
	"context"

	"github.com/alexmourapb/url-shortener/app/domain/shortener/entities"
)

type Repository interface {
	Save(ctx context.Context, id, url string) error
	GetByID(ctx context.Context, id string) (*entities.ShortURL, error)
}

type CacheRepository interface {
	Save(key string, value interface{}) error
	GetURL(key string) (*entities.ShortURL, error)
}
