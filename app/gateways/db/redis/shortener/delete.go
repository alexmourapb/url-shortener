package shortener

func (r Repository) Delete(key string) error {
	redisConn := r.redisPool.Get()
	defer redisConn.Close()

	_, err := redisConn.Do(argDel, key)
	if err != nil {
		return err
	}

	return nil
}
