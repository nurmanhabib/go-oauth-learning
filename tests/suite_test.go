package tests_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/nurmanhabib/go-oauth2/config"
	"gitlab.com/nurmanhabib/go-oauth2/tests"
)

func TestNewApplicationRepository(t *testing.T) {
	t.Run("if plain suite test", func(t *testing.T) {
		suite := tests.New()

		assert.True(t, suite.Config.TestMode)
	})

	t.Run("if some connection given", func(t *testing.T) {
		suite := tests.New(
			tests.WithDatabase(),
			tests.WithRedis(),
		)

		assert.NotEqual(t, config.Database{}, suite.Config.Database)
		assert.NotEqual(t, config.Redis{}, suite.Config.Redis)

		assert.NotNil(t, suite.DB)
		assert.NotNil(t, suite.Redis)
	})
}
