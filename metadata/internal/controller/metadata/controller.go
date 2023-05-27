package metadata

import (
	"context"
	"errors"

	"movieexample.com/metadata/internal/repository"
	"movieexample.com/metadata/pkg/model"
)

var ErrorNotFound = errors.New("not found")

type MetadataRepository interface {
	Get(ctx context.Context, id string) (*model.Metadata, error)
	Put(ctx context.Context, id string, metadata *model.Metadata) error
}

type Controller struct {
	repo MetadataRepository
}

func New(repo MetadataRepository) *Controller {
	return &Controller{
		repo: repo,
	}
}

func (c *Controller) Get(ctx context.Context, id string) (*model.Metadata, error) {
	m, err := c.repo.Get(ctx, id)
	if err != nil {
		if err == repository.ErrorNotFound {
			return nil, ErrorNotFound
		}

		return nil, err
	}

	return m, nil
}
