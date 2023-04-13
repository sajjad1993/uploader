package merger

import (
	"OMPFinex-CodeChallenge/internal/contract/chunk"
	"OMPFinex-CodeChallenge/internal/contract/image"
	"OMPFinex-CodeChallenge/pkg/log"
)

type Merger struct {
	//todo proto adapter
	imageRepo image.Repository
	chunkRepo chunk.Reader
	logger    log.Logger
}

func New(imageRepo image.Repository, chunkRepo chunk.Repository, logger log.Logger) UseCase {
	return &Merger{
		imageRepo: imageRepo,
		chunkRepo: chunkRepo,
		logger:    logger,
	}
}
