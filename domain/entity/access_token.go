package entity

import (
	"database/sql"
	"time"
)

// AccessToken is an access token entity container.
type AccessToken struct {
	ID              int
	ResourceOwnerID int
	ApplicationID   int
	Token           string
	RefreshToken    string
	ExpiresIn       int
	Scopes          string
	RevokedAt       sql.NullTime
	CreatedAt       time.Time
}
