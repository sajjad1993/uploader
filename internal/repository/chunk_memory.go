package repository

import (
	"OMPFinex-CodeChallenge/internal/contract/chunk"
	"OMPFinex-CodeChallenge/internal/model"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

// ChunkRepo postgres repo
type ChunkMemory struct {
}

// NewChunkRepo creates new repository
func NewChunkMemory(dbPool *pgxpool.Pool) chunk.Repository {
	return &ChunkRepo{}
}

func (c ChunkMemory) GetAll(ctx context.Context, sha string) ([]model.Chunk, error) {
	return nil, nil
}

func (c ChunkMemory) Get(ctx context.Context, sha string, id int) (*model.Chunk, error) {
	return nil, nil
}

func (c ChunkMemory) Save(ctx context.Context, chunk model.Chunk) error {
	return nil
}
