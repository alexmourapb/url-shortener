package usecase

import (
	"context"

	"github.com/rs/zerolog"

	"github.com/alexmourapb/url-shortener/app/common/domain"
	"github.com/alexmourapb/url-shortener/app/common/shared"
)

func (s Shortener) Get(ctx context.Context, log *zerolog.Logger, id string) (string, error) {
	const operation = "UseCase.Get"
	var output string
	cacheOutput, err := s.cache.GetURL(id)
	if err != nil {
		log.Err(err).Msg("failed to get data from cache")
	}

	if cacheOutput != nil {
		output = cacheOutput.URL
	} else {
		dbOtput, err := s.db.GetByID(ctx, id)
		if err != nil {
			return "", domain.Error(operation, err)
		}

		if !dbOtput.Active {
			return "", shared.ErrURLNotFound
		}

		go func() {
			if err := s.cache.Save(id, dbOtput); err != nil {
				log.Err(err).Msg("failed to save data in to cache")
			}
		}()

		output = dbOtput.URL
	}

	return output, nil
}
