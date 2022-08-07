package userprovider

import (
	"context"
	"fmt"
	"strconv"

	"gitlab.com/nurmanhabib/go-oauth2/domain/entity"
	"gitlab.com/nurmanhabib/go-oauth2/domain/repository"
	"gitlab.com/nurmanhabib/go-oauth2/pkg/auth"
	"gitlab.com/nurmanhabib/go-oauth2/pkg/hashing"
	"gorm.io/gorm"
)

// DatabaseUserProvider is a container for the needs of the user provider in the database.
type DatabaseUserProvider struct {
	repo   repository.UserRepository
	hasher hashing.Hasher
}

// FindByIdentifier is a function to search for users by identifier.
func (p *DatabaseUserProvider) FindByIdentifier(ctx context.Context, identifier string) (auth.User, error) {
	id, _ := strconv.Atoi(identifier)
	userEntity, err := p.repo.Find(ctx, &entity.User{ID: id})
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return nil, auth.ErrUserNotFound
		default:
			return nil, fmt.Errorf("error find by credentials: %w", err)
		}
	}

	return userEntity, nil
}

// FindByCredentials is a function to search for users based on credentials.
func (p *DatabaseUserProvider) FindByCredentials(ctx context.Context, credentials auth.Credentials) (auth.User, error) {
	userEntity, err := p.repo.FindByCredentials(ctx, credentials.WithoutPassword())
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return nil, auth.ErrUserNotFound
		default:
			return nil, fmt.Errorf("error find by credentials: %w", err)
		}
	}

	return userEntity, nil
}

// ValidateCredentials is a function for authentication by comparing the password credentials.
func (p *DatabaseUserProvider) ValidateCredentials(ctx context.Context, user auth.User, credentials auth.Credentials) error {
	ok := p.hasher.Check(ctx, credentials.GetPassword(), user.GetAuthPassword())
	if !ok {
		return auth.ErrInvalidCredentials
	}

	return nil
}
