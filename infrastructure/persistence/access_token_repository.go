package persistence

import (
	"context"

	"gitlab.com/nurmanhabib/go-oauth2/domain/entity"
	"gorm.io/gorm"
)

// AccessTokenRepository is a container for implementing repositories in a persistent way to the database.
type AccessTokenRepository struct {
	db *gorm.DB
}

// NewAccessTokenRepository is a concrete repository constructor.
func NewAccessTokenRepository(db *gorm.DB) *AccessTokenRepository {
	return &AccessTokenRepository{db: db}
}

// Find is a function to get access token from within the database.
func (a *AccessTokenRepository) Find(ctx context.Context, token *entity.AccessToken) (*entity.AccessToken, error) {
	// TODO implement me
	panic("implement me")
}

// Save is a function to store access token into the database.
func (a *AccessTokenRepository) Save(ctx context.Context, token *entity.AccessToken) error {
	// TODO implement me
	panic("implement me")
}
