package tests

import (
	"log"

	"gitlab.com/nurmanhabib/go-oauth2/config"
	"gitlab.com/nurmanhabib/go-oauth2/infrastructure/provider/connection"
)

// WithRedis is a function to establish a redis connection in the test suite.
func WithRedis() SuiteOption {
	return SuiteOption{
		Config: func(conf *config.Config) {
			config.WithRedis()(conf)
		},
		Suite: func(suite *Suite) {
			rdb, errRDB := connection.NewRedisConnection(suite.Config)
			if errRDB != nil {
				log.Fatalf("Unable to connect redis: %v", errRDB)
			}

			suite.Redis = rdb
		},
	}
}
