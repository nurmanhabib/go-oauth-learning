package connection

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v9"
	"gitlab.com/nurmanhabib/go-oauth2/config"
)

// NewRedisConnection is a function to establish a redis connection.
func NewRedisConnection(config *config.Config) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Redis.Host, config.Redis.Port),
		Password: config.Redis.Password,
		DB:       config.Redis.DB,
	})

	err := rdb.Ping(context.Background()).Err()
	if err != nil {
		return nil, err
	}

	return rdb, nil
}
