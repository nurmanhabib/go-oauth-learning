package repository

import (
	"context"

	"gitlab.com/nurmanhabib/go-oauth2/domain/entity"
)

// AccessGrantRepository is a repository interface for access grant entities.
type AccessGrantRepository interface {
	Find(ctx context.Context, grant *entity.AccessGrant) (*entity.AccessGrant, error)
	Save(ctx context.Context, grant *entity.AccessGrant) error
}
