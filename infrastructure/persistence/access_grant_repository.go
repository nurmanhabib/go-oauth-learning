package persistence

import (
	"context"

	"gitlab.com/nurmanhabib/go-oauth2/domain/entity"
	"gorm.io/gorm"
)

// AccessGrantRepository is a container for implementing repositories in a persistent way to the database.
type AccessGrantRepository struct {
	db *gorm.DB
}

// NewAccessGrantRepository is a concrete repository constructor.
func NewAccessGrantRepository(db *gorm.DB) *AccessGrantRepository {
	return &AccessGrantRepository{db: db}
}

// Find is a function to get access grant from within the database.
func (a *AccessGrantRepository) Find(ctx context.Context, grant *entity.AccessGrant) (*entity.AccessGrant, error) {
	// TODO implement me
	panic("implement me")
}

// Save is a function to store access grant into the database.
func (a *AccessGrantRepository) Save(ctx context.Context, grant *entity.AccessGrant) error {
	// TODO implement me
	panic("implement me")
}
