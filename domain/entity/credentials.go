package entity

import (
	"time"

	"gorm.io/gorm"
)

// Credentials is an application credentials entity container.
type Credentials struct {
	ID           string
	ClientID     string
	ClientSecret string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}
