package shortener

import (
	"context"

	"github.com/rs/zerolog"

	"github.com/alexmourapb/url-shortener/app/domain/shortener/vos"
)

type UseCase interface {
	Create(ctx context.Context, url string) (string, error)
	Get(ctx context.Context, log *zerolog.Logger, id string) (string, error)
	Update(ctx context.Context, log *zerolog.Logger, input vos.UpdateInput) error
}
