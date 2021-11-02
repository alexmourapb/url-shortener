package usecase

import (
	"github.com/alexmourapb/url-shortener/app/domain/shortener"
)

var _ shortener.UseCase = Shortener{}

type Shortener struct {
	db    shortener.Repository
	cache shortener.CacheRepository
}

func NewShortenerUseCase(
	db shortener.Repository,
	cache shortener.CacheRepository,
) *Shortener {
	return &Shortener{
		db:    db,
		cache: cache,
	}
}
