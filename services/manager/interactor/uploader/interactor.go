package interactor

import (
	"OMPFinex-CodeChallenge/internal/contract/chunk"
	"OMPFinex-CodeChallenge/internal/contract/image"
	"OMPFinex-CodeChallenge/pkg/log"
)

type Manager struct {
	imageRepo       image.Repository
	chunkRepo       chunk.Repository
	logger          log.Logger
	imageChannel    chan string
	imageCollectors map[string]*ImageCollector
}

func New(imageRepo image.Repository, chunkRepo chunk.Repository, logger log.Logger) *Manager {
	channel := make(chan string)
	return &Manager{
		imageRepo:    imageRepo,
		chunkRepo:    chunkRepo,
		logger:       logger,
		imageChannel: channel,
	}
}
