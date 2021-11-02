package entities

import "time"

type ShortURL struct {
	ID        string
	URL       string
	CreatedAt time.Time
}
