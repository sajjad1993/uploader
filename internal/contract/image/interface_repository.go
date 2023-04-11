package image

import (
	"OMPFinex-CodeChallenge/internal/model"
	"context"
)

type Reader interface {
	CheckDuplicate(ctx context.Context, image model.Image) error
	Get(ctx context.Context, sha string) (*model.Image, error)
}

type Writer interface {
	Save(ctx context.Context, image *model.Image)
}

type Repository interface {
	Reader
	Writer
}
