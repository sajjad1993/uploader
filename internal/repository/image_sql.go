package repository

import (
	"OMPFinex-CodeChallenge/internal/contract/image"
	"OMPFinex-CodeChallenge/internal/model"
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

// ImageRepo postgres repo
type ImageRepo struct {
	pgxPool *pgxpool.Pool
}

// NewImageRepo creates new repository
func NewImageRepo(dbPool *pgxpool.Pool) image.Repository {
	return &ImageRepo{
		pgxPool: dbPool,
	}
}

// todo delete this
func (i ImageRepo) CheckDuplicate(ctx context.Context, image model.Image) (bool, error) {
	sqlStatement := `
	select count(*) from images where sha = $1`

	q := i.pgxPool.QueryRow(ctx, sqlStatement, image.Sha256)
	var totalCount int64

	err := q.Scan(&totalCount)
	if err != nil {
		return false, err
	}
	return totalCount > 0, nil
}

func (i ImageRepo) DoesExist(ctx context.Context, sha string) (bool, error) {
	sqlStatement := `
	select count(*) from images where sha = $1`

	q := i.pgxPool.QueryRow(ctx, sqlStatement, sha)
	var totalCount int64

	err := q.Scan(&totalCount)
	if err != nil {
		return false, err
	}
	return totalCount > 0, nil
}

func (i ImageRepo) Get(ctx context.Context, sha string) (*model.Image, error) {
	sqlStatement := `
	select sha , size, chunk_size , "status" , "created_at" from  images where sha = $1 limit 1`
	q := i.pgxPool.QueryRow(ctx, sqlStatement, sha)
	return parseImage(q)
}

func (i ImageRepo) Save(ctx context.Context, image model.Image) error {

	sqlStatement := `
	INSERT INTO images (sha , size, chunk_size , "status" , "created_at")
			VALUES ($1, $2, $3, $4,$5)`

	_, err := i.pgxPool.Exec(ctx, sqlStatement, image.Sha256, image.Size, image.ChunkSize, image.Status, time.Now().UTC())
	return err
}

func (i ImageRepo) Update(ctx context.Context, image model.Image) error {
	sqlStatement := `
	update images set  "status" = $1 , "data" = $2  where sha = $3`
	_, err := i.pgxPool.Exec(ctx, sqlStatement, image.Status, image.Data, image.Sha256)
	return err
}

func parseImage(q pgx.Row) (*model.Image, error) {
	resp := model.Image{}
	err := q.Scan(&resp.Sha256, &resp.Size, &resp.ChunkSize, &resp.Status, &resp.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
