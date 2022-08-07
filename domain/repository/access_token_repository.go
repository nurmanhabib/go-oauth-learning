package repository

import (
	"context"

	"gitlab.com/nurmanhabib/go-oauth2/domain/entity"
)

// AccessTokenRepository is a repository interface for access token entities.
type AccessTokenRepository interface {
	Find(ctx context.Context, grant *entity.AccessToken) (*entity.AccessToken, error)
	Save(ctx context.Context, grant *entity.AccessToken) error
}
