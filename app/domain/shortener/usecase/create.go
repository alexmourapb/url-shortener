package usecase

import (
	"context"
	"math/rand"

	"github.com/alexmourapb/url-shortener/app/common/domain"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func (s Shortener) Create(ctx context.Context, url string) (string, error) {
	const operation = "UseCase.Create"
	id := shortGenerate()
	err := s.db.Save(ctx, id, url)
	if err != nil {
		return "", domain.Error(operation, err)
	}

	return id, nil
}

func shortGenerate() string {
	var chars = []rune(alphabet)
	output := make([]rune, 8)
	for i := range output {
		output[i] = chars[rand.Intn(len(chars))]
	}

	return string(output)
}
