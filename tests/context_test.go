package tests_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/nurmanhabib/go-oauth2/tests"
)

type contextKey string

func TestWithContext(t *testing.T) {
	t.Run("if no custom context", func(t *testing.T) {
		suite := tests.New()

		assert.Equal(t, context.Background(), suite.Ctx)
	})

	t.Run("if with value context", func(t *testing.T) {
		keyData := contextKey("data")
		ctx := context.Background()
		ctx = context.WithValue(ctx, keyData, 123)

		suite := tests.New(
			tests.WithContext(ctx),
		)

		assert.Equal(t, ctx, suite.Ctx)
		assert.Equal(t, 123, suite.Ctx.Value(keyData))
	})
}
