package uploader

import (
	"OMPFinex-CodeChallenge/internal/contract/image"
	"OMPFinex-CodeChallenge/pkg/log"
	"time"
)

type Uploader struct {
	imageRepo image.Reader
	logger    log.Logger
	retry     uint
	sleepTime time.Duration
}

func New(imageRepo image.Repository, retry uint, sleepTime time.Duration, logger log.Logger) UseCase {
	return &Uploader{
		imageRepo: imageRepo,
		logger:    logger,
		retry:     retry,
		sleepTime: sleepTime,
	}
}
