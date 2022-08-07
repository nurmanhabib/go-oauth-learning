package repository

import (
	"context"

	"gitlab.com/nurmanhabib/go-oauth2/domain/entity"
	"gitlab.com/nurmanhabib/go-oauth2/pkg/auth"
)

// UserRepository is a repository interface for user entities.
type UserRepository interface {
	Find(ctx context.Context, user *entity.User) (*entity.User, error)
	FindByCredentials(ctx context.Context, credentials auth.Credentials) (*entity.User, error)
	Save(ctx context.Context, user *entity.User) error
}
