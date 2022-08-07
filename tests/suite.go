package tests

import (
	"context"

	"github.com/go-redis/redis/v9"
	"gitlab.com/nurmanhabib/go-oauth2/config"
	"gorm.io/gorm"
)

// Suite is a container for test needs.
type Suite struct {
	Config *config.Config
	Ctx    context.Context

	DB    *gorm.DB
	Redis *redis.Client
}

// New is a new test constructor.
func New(options ...SuiteOption) *Suite {
	// Required Load .env.testing
	loadEnvTesting()

	ctx := context.Background()
	conf := config.NewTestMode()

	// Apply Config Option
	for _, apply := range options {
		if apply.Config != nil {
			apply.Config(conf)
		}
	}

	// Initialize New Suite
	suite := &Suite{
		Config: conf,
		Ctx:    ctx,
	}

	// Apply Suite Option
	for _, apply := range options {
		if apply.Suite != nil {
			apply.Suite(suite)
		}
	}

	return suite
}
