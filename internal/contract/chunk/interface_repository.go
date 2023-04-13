package chunk

import (
	"OMPFinex-CodeChallenge/internal/model"
	"context"
)

type Reader interface {
	GetAll(ctx context.Context, sha string) ([]model.Chunk, error)
	Get(ctx context.Context, sha string, id int) (*model.Chunk, error)
}

type Writer interface {
	Save(ctx context.Context, chunk model.Chunk) error
}

type Repository interface {
	Reader
	Writer
}
