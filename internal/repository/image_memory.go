package repository

import (
	"OMPFinex-CodeChallenge/internal/contract/image"
	"OMPFinex-CodeChallenge/internal/model"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

// ImageMemory  memory repo
type ImageMemory struct {
}

// NewImageMemory  creates new repository
func NewImageMemory(dbPool *pgxpool.Pool) image.Repository {
	return &ImageMemory{}
}

func (i ImageMemory) DoesExist(ctx context.Context, sha string) (bool, error) {
	return false, nil
}

func (i ImageMemory) Get(ctx context.Context, sha string) (*model.Image, error) {
	return nil, nil
}

func (i ImageMemory) Save(ctx context.Context, image model.Image) error {

	return nil
}

func (i ImageMemory) Update(ctx context.Context, image model.Image) error {
	return nil
}
