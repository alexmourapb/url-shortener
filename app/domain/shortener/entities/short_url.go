package entities

import "time"

type ShortURL struct {
	ID        string
	URL       string
	Active    bool
	CreatedAt time.Time
}
