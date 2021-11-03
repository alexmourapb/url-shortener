package usecase

import (
	"context"
	"math/rand"
	"time"

	"github.com/alexmourapb/url-shortener/app/common/domain"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func (s Shortener) Create(ctx context.Context, url string) (string, error) {
	const operation = "UseCase.Create"
	id := shortKeyGenerate()
	err := s.db.Save(ctx, id, url)
	if err != nil {
		return "", domain.Error(operation, err)
	}

	return id, nil
}

// Generate a 63-bit random string
func shortKeyGenerate() string {
	b := make([]byte, 7)

	currentTime := time.Now().UnixNano()
	rand.Seed(currentTime)

	for i := range b {
		b[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(b)
}
