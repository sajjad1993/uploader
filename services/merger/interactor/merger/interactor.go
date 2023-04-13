package merger

import (
	"OMPFinex-CodeChallenge/internal/contract/chunk"
	"OMPFinex-CodeChallenge/internal/contract/image"
	"OMPFinex-CodeChallenge/pkg/log"
	"time"
)

type Merger struct {
	//todo proto adapter
	imageRepo image.Repository
	chunkRepo chunk.Reader
	logger    log.Logger
	timeout   time.Duration
}

func New(imageRepo image.Repository, chunkRepo chunk.Repository, logger log.Logger, timeout time.Duration) UseCase {
	return &Merger{
		imageRepo: imageRepo,
		chunkRepo: chunkRepo,
		logger:    logger,
		timeout:   timeout,
	}
}
