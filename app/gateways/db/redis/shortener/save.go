package shortener

import (
	"encoding/json"
)

func (r Repository) Save(key string, value interface{}) error {
	redisConn := r.redisPool.Get()
	defer redisConn.Close()

	structBytes, err := json.Marshal(value)
	if err != nil {
		return err
	}

	_, err = redisConn.Do(argSet, key, string(structBytes), argEx, ExpireKeyTime, argNx)
	if err != nil {
		return err
	}

	return nil
}
