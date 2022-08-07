package entity

import (
	"database/sql"
	"time"
)

// AccessGrant Application is an access grant entity container.
type AccessGrant struct {
	ID                  int
	ResourceOwnerID     int
	ApplicationID       int
	Token               string
	RedirectURI         string
	ExpiresIn           int
	Scopes              string
	CodeChallenge       string
	CodeChallengeMethod string
	RevokedAt           sql.NullTime
	CreatedAt           time.Time
}
