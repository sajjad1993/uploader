package uploader

import (
	"OMPFinex-CodeChallenge/internal/contract/chunk"
	"OMPFinex-CodeChallenge/internal/contract/image"
	"OMPFinex-CodeChallenge/pkg/log"
)

type Uploader struct {
	imageRepo image.Reader
	logger    log.Logger
}

func New(imageRepo image.Repository, chunkRepo chunk.Repository, logger log.Logger) UseCase {
	return &Uploader{
		imageRepo: imageRepo,
		logger:    logger,
	}
}
