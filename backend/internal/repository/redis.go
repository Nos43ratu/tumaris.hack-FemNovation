package repository

import (
	"github.com/go-redis/redis/v7"
	"go.uber.org/zap"
)

func RedisConn(sugar *zap.SugaredLogger) (*redis.Client, error) {
	username := ""
	password := ""
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6666",
		Password: password,
		Username: username,
		DB:       0,
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		return nil, err
	}

	return rdb, nil
}
