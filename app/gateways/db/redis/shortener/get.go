package shortener

import (
	"encoding/json"

	"github.com/alexmourapb/url-shortener/app/domain/shortener/entities"
)

func (r Repository) GetURL(key string) (*entities.ShortURL, error) {
	redisConn := r.redisPool.Get()
	defer redisConn.Close()

	value, err := redisConn.Do(argGet, key)
	if err != nil {
		return nil, err
	}

	if value == nil {
		return nil, nil
	}

	var output entities.ShortURL

	err = json.Unmarshal(value.([]byte), &output)
	if err != nil {
		return nil, err
	}

	return &output, nil
}
