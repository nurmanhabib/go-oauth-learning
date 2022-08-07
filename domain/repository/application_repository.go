package repository

import (
	"context"

	"gitlab.com/nurmanhabib/go-oauth2/domain/entity"
)

// ApplicationRepository is a repository interface for application entities.
type ApplicationRepository interface {
	Find(ctx context.Context, application *entity.Application) (*entity.Application, error)
	Save(ctx context.Context, application *entity.Application) error
}
