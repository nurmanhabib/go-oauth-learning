package tests

import (
	"log"

	"gitlab.com/nurmanhabib/go-oauth2/config"
	"gitlab.com/nurmanhabib/go-oauth2/infrastructure/provider/connection"
)

// WithDatabase is a function to establish a database connection in the test suite.
func WithDatabase() SuiteOption {
	return SuiteOption{
		Config: func(conf *config.Config) {
			config.WithDatabase()(conf)
		},
		Suite: func(suite *Suite) {
			db, errDB := connection.NewDBConnection(suite.Config)
			if errDB != nil {
				log.Fatalf("Unable to connect database: %v", errDB)
			}

			suite.DB = db
		},
	}
}
