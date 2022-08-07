package tests

import (
	"context"
)

// WithContext is a function to set a custom context into the test suite.
func WithContext(ctx context.Context) SuiteOption {
	return SuiteOption{
		Suite: func(suite *Suite) {
			suite.Ctx = ctx
		},
	}
}
