package tests

import "gitlab.com/nurmanhabib/go-oauth2/config"

// Option is an interface for applying test suite options.
type Option func(*Suite)

// SuiteOption is a struct for apply config and test suite.
type SuiteOption struct {
	Config config.Option
	Suite  Option
}
