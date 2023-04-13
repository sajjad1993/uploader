package repository

import (
	"OMPFinex-CodeChallenge/internal/contract/chunk"
	"OMPFinex-CodeChallenge/internal/model"
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

// ChunkRepo postgres repo
type ChunkRepo struct {
	pgxPool *pgxpool.Pool
}

// NewChunkRepo creates new repository
func NewChunkRepo(dbPool *pgxpool.Pool) chunk.Repository {
	return &ChunkRepo{
		pgxPool: dbPool,
	}
}

func (c ChunkRepo) GetAll(ctx context.Context, sha string) ([]model.Chunk, error) {
	sqlStatement := `
	select "id" , "size", "data" , "sha" , "created_at" from  chunks where sha = $1 ORDER BY id ASC`
	q, err := c.pgxPool.Query(ctx, sqlStatement, sha)
	if err != nil {
		return nil, err
	}
	return parseManyChunk(q)
}

func (c ChunkRepo) Get(ctx context.Context, sha string, id int) (*model.Chunk, error) {
	sqlStatement := `
	select "id" , "size", "data" , "sha" , "created_at" from  chunks where sha = $1 AND id = $2 limit 1 `
	q := c.pgxPool.QueryRow(ctx, sqlStatement, sha, id)
	return parseChunk(q)
}

func (c ChunkRepo) Save(ctx context.Context, chunk model.Chunk) error {
	sqlStatement := `
	INSERT INTO chunks ("id" , "size", "data" , "sha" , "created_at")
			VALUES ($1, $2, $3, $4,$5)`
	_, err := c.pgxPool.Exec(ctx, sqlStatement, chunk.ID, chunk.Size, chunk.Data, chunk.Sha256, time.Now().UTC())
	return err
}

func parseChunk(q pgx.Row) (*model.Chunk, error) {
	resp := model.Chunk{}
	err := q.Scan(&resp.ID, &resp.Size, &resp.Data, &resp.Sha256, &resp.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func parseManyChunk(q pgx.Rows) ([]model.Chunk, error) {
	var resp []model.Chunk

	for q.Next() {
		tmp, err := parseChunk(q)
		if err != nil {
			return nil, err
		}

		resp = append(resp, *tmp)
	}

	return resp, nil
}
