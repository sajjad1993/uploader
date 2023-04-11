package interactor

import (
	"OMPFinex-CodeChallenge/internal/contract/chunk"
	"OMPFinex-CodeChallenge/internal/contract/image"
	"OMPFinex-CodeChallenge/pkg/log"
	"OMPFinex-CodeChallenge/services/manager/entity"
)

type Manager struct {
	imageRepo       image.Repository
	chunkRepo       chunk.Repository
	logger          log.Logger
	imageChannel    chan entity.Image
	imageCollectors map[string]ImageCollector
}

func New(imageRepo image.Repository, chunkRepo chunk.Repository, logger log.Logger) *Manager {
	channel := make(chan entity.Image)
	return &Manager{
		imageRepo:    imageRepo,
		chunkRepo:    chunkRepo,
		logger:       logger,
		ImageChannel: channel,
	}
}
