package entity

import (
	"time"

	"gorm.io/gorm"
)

// ApplicationVersion is a type to define the version of the application.
type ApplicationVersion string

// ApplicationVX is the application version constant.
const (
	ApplicationV1 = ApplicationVersion("v1")
	ApplicationV2 = ApplicationVersion("v2")
)

// Application is an application entity container.
type Application struct {
	ID          string
	Name        string
	Version     ApplicationVersion
	Credentials Credentials
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

// IsV1 is a function to ensure the application version 1.
func (a *Application) IsV1() bool {
	return a.IsVersion(ApplicationV1)
}

// IsV2 is a function to ensure the application version 2.
func (a *Application) IsV2() bool {
	return a.IsVersion(ApplicationV2)
}

// IsVersion is a function to confirm the version of the application.
func (a *Application) IsVersion(version ApplicationVersion) bool {
	if a == nil {
		return false
	}

	return a.Version == version
}
