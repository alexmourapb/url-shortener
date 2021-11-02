package usecase

import (
	"context"

	"github.com/alexmourapb/url-shortener/app/common/domain"
)

func (s Shortener) Get(ctx context.Context, id string) (string, error) {
	const operation = "UseCase.Get"
	var output string
	cacheOutput, err := s.cache.GetURL(id)
	if err != nil {
		return "", domain.Error(operation, err)
	}

	if cacheOutput != nil {
		output = cacheOutput.URL
	} else {
		odbOtput, err := s.db.GetByID(ctx, id)
		if err != nil {
			return "", domain.Error(operation, err)
		}
		go s.cache.Save(id, odbOtput)
		output = odbOtput.URL
	}

	return output, nil
}
