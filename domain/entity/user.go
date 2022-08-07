package entity

import (
	"time"

	"gorm.io/gorm"
)

// User is a user entity container.
type User struct {
	ID        string
	Name      string
	PrivyID   string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

// GetAuthIdentifier is a function to take as auth identifier.
func (u *User) GetAuthIdentifier() string {
	if u == nil {
		return ""
	}

	return u.ID
}

// GetAuthPassword is a function to take as auth password.
func (u *User) GetAuthPassword() string {
	if u == nil {
		return ""
	}

	return u.Password
}
