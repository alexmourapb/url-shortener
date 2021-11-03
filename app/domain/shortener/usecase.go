package shortener

import (
	"context"

	"github.com/rs/zerolog"
)

type UseCase interface {
	Create(ctx context.Context, url string) (string, error)
	Get(ctx context.Context, log *zerolog.Logger, id string) (string, error)
}
