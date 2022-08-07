package tests_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/nurmanhabib/go-oauth2/config"
	"gitlab.com/nurmanhabib/go-oauth2/tests"
)

func TestWithRedis(t *testing.T) {
	t.Run("if no redis", func(t *testing.T) {
		suite := tests.New()

		assert.Equal(t, config.Redis{}, suite.Config.Redis)
		assert.Nil(t, suite.Redis)
	})

	t.Run("if with redis", func(t *testing.T) {
		suite := tests.New(
			tests.WithRedis(),
		)

		assert.NotEqual(t, config.Redis{}, suite.Config.Redis)
		assert.NotNil(t, suite.Redis)
	})
}
