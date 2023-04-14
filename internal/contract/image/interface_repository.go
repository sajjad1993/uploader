package image

import (
	"OMPFinex-CodeChallenge/internal/model"
	"context"
)

type Reader interface {
	DoesExist(ctx context.Context, sha string) (bool, error)
	Get(ctx context.Context, sha string) (*model.Image, error)
}

type Writer interface {
	Save(ctx context.Context, image model.Image) error
	Update(ctx context.Context, image model.Image) error
}

type Repository interface {
	Reader
	Writer
}
