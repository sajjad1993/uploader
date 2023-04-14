package repository

import (
	"OMPFinex-CodeChallenge/internal/contract/image"
	"OMPFinex-CodeChallenge/internal/model"
	"OMPFinex-CodeChallenge/pkg/errs"
	"context"
	"fmt"
	"path/filepath"
)

// ImageMemory  memory repo
type ImageMemory struct {
	imageDir string
	chunkDir string
}

// NewImageMemory  creates new repository
func NewImageMemory(imageDir, chunkDir string) (image.Repository, error) {
	err := createDir(imageDir)
	if err != nil {
		return nil, err
	}
	err = createDir(chunkDir)
	if err != nil {
		return nil, err
	}
	return &ImageMemory{
		imageDir: imageDir,
		chunkDir: chunkDir,
	}, nil
}

func (i ImageMemory) DoesExist(ctx context.Context, sha string) (bool, error) {
	return isDirExists(filepath.Join(i.chunkDir, sha))
}

func (i ImageMemory) Get(ctx context.Context, sha string) (*model.Image, error) {

	ok, err := isDirExists(filepath.Join(i.chunkDir, sha))
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errs.NewNotFoundError(fmt.Sprintf("sha %s doesn't exist", sha))
	}
	imageModel := model.Image{
		Sha256: sha,
		Status: string(model.UnCompletedStatus),
	}
	ok, err = isFileExists(filepath.Join(i.imageDir, sha))
	if err != nil {
		return nil, err
	}
	if !ok {
		return &imageModel, nil
	}
	imageModel.Data = filepath.Join(i.imageDir, sha)
	imageModel.Status = string(model.ReadyStatus)
	return &imageModel, nil

}

func (i ImageMemory) Save(ctx context.Context, image model.Image) error {
	return nil
}

func (i ImageMemory) Update(ctx context.Context, image model.Image) error {
	return nil
}
