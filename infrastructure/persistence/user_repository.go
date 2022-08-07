package persistence

import (
	"context"

	"gitlab.com/nurmanhabib/go-oauth2/domain/entity"
	"gitlab.com/nurmanhabib/go-oauth2/pkg/auth"
	"gorm.io/gorm"
)

// UserRepository is a container for implementing repositories in a persistent way to the database.
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository is a concrete repository constructor.
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Find is a function to get user from within the database.
func (u *UserRepository) Find(ctx context.Context, user *entity.User) (*entity.User, error) {
	// TODO implement me
	panic("implement me")
}

// FindByCredentials is a function to get user by auth credentials from within the database.
func (u *UserRepository) FindByCredentials(ctx context.Context, credentials auth.Credentials) (*entity.User, error) {
	// TODO implement me
	panic("implement me")
}

// Save is a function to store user into the database.
func (u *UserRepository) Save(ctx context.Context, user *entity.User) error {
	// TODO implement me
	panic("implement me")
}
