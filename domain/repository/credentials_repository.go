package repository

import (
	"context"

	"gitlab.com/nurmanhabib/go-oauth2/domain/entity"
)

// CredentialsRepository is a repository interface for credentials entities.
type CredentialsRepository interface {
	Find(ctx context.Context, credentials *entity.Credentials) (*entity.Credentials, error)
	Save(ctx context.Context, credentials *entity.Credentials) error
}
