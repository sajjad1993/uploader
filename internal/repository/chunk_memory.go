package repository

import (
	"OMPFinex-CodeChallenge/internal/contract/chunk"
	"OMPFinex-CodeChallenge/internal/model"
	"OMPFinex-CodeChallenge/pkg/errs"
	"context"
	"fmt"
	"path/filepath"
	"sort"
	"strconv"
)

// ChunkMemory  postgres repo
type ChunkMemory struct {
	imageDir string
	chunkDir string
}

// NewChunkMemory  creates new repository
func NewChunkMemory(imageDir, chunkDir string) (chunk.Repository, error) {
	err := createDir(imageDir)
	if err != nil {
		return nil, err
	}
	err = createDir(chunkDir)
	if err != nil {
		return nil, err
	}
	return &ChunkMemory{
		imageDir: imageDir,
		chunkDir: chunkDir,
	}, nil
}

func (c ChunkMemory) GetAll(ctx context.Context, sha string) ([]model.Chunk, error) {
	imageChunksDir := filepath.Join(c.chunkDir, sha)
	ok, err := isDirExists(imageChunksDir)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errs.NewNotFoundError(fmt.Sprintf("sha %s doesn't exist", sha))
	}
	var chunks []model.Chunk
	contents, err := listDirFiles(imageChunksDir)
	if err != nil {
		return nil, err
	}
	for _, content := range contents {
		id, err := strconv.Atoi(content)

		if err != nil {
			return nil, err
		}
		c := model.Chunk{
			Sha256: sha,
			ID:     uint(id),
			Data:   filepath.Join(imageChunksDir, content),
		}
		chunks = append(chunks, c)
	}
	sort.Slice(chunks, func(i, j int) bool {
		return chunks[i].ID < chunks[j].ID
	})
	return chunks, nil
}

func (c ChunkMemory) Get(ctx context.Context, sha string, id int) (*model.Chunk, error) {
	return nil, nil
}

func (c ChunkMemory) Save(ctx context.Context, chunk model.Chunk) error {
	return nil
}
