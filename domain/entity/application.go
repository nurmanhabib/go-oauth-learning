package entity

import (
	"time"
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
	ID            int
	Name          string
	ClientID      string             `gorm:"column:uid"`
	ClientSecret  string             `gorm:"column:secret"`
	Version       ApplicationVersion `gorm:"-"`
	SuperApp      bool               `gorm:"column:super_app"`
	RedirectURI   string             `gorm:"column:redirect_uri"`
	Scopes        string
	Confidential  bool
	SingleSession bool
	Status        bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
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
