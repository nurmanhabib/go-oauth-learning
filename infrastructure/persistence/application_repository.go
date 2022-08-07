package persistence

import (
	"context"

	"gitlab.com/nurmanhabib/go-oauth2/domain/entity"
	"gorm.io/gorm"
)

// ApplicationRepository is a container for implementing repositories in a persistent way to the database.
type ApplicationRepository struct {
	db *gorm.DB
}

// NewApplicationRepository is a concrete repository constructor.
func NewApplicationRepository(db *gorm.DB) *ApplicationRepository {
	return &ApplicationRepository{db: db}
}

// Find is a function to get application from within the database.
func (a *ApplicationRepository) Find(ctx context.Context, application *entity.Application) (*entity.Application, error) {
	// TODO implement me
	panic("implement me")
}

// Save is a function to store application into the database.
func (a *ApplicationRepository) Save(ctx context.Context, application *entity.Application) error {
	// TODO implement me
	panic("implement me")
}
