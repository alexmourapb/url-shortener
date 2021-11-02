package shortener

import (
	"context"
)

type UseCase interface {
	Create(ctx context.Context, url string) (string, error)
	Get(ctx context.Context, id string) (string, error)
}
