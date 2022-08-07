package tests_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/nurmanhabib/go-oauth2/config"
	"gitlab.com/nurmanhabib/go-oauth2/tests"
)

func TestWithDatabase(t *testing.T) {
	t.Run("if no database", func(t *testing.T) {
		suite := tests.New()

		assert.Equal(t, config.Database{}, suite.Config.Database)
		assert.Nil(t, suite.DB)
	})

	t.Run("if with database", func(t *testing.T) {
		suite := tests.New(
			tests.WithDatabase(),
		)

		assert.NotEqual(t, config.Database{}, suite.Config.Database)
		assert.NotNil(t, suite.DB)
	})
}
