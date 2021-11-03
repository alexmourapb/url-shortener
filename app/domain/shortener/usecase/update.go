package usecase

import (
	"context"

	"github.com/rs/zerolog"

	"github.com/alexmourapb/url-shortener/app/common/domain"
	"github.com/alexmourapb/url-shortener/app/domain/shortener/vos"
)

func (s Shortener) Update(ctx context.Context, log *zerolog.Logger, input vos.UpdateInput) error {
	const operation = "UseCase.Update"
	dbOutput, err := s.db.GetByID(ctx, input.ID)
	if err != nil {
		return domain.Error(operation, err)
	}

	if input.URL != nil {
		dbOutput.URL = *input.URL
	}

	if input.Active != nil {
		dbOutput.Active = *input.Active
	}

	err = s.db.Update(ctx, *dbOutput)
	if err != nil {
		return domain.Error(operation, err)
	}
	go func() {
		if err := s.cache.Delete(input.ID); err != nil {
			log.Err(err).Msg("failed to delete data in to cache")
		}
	}()

	return nil
}
