package donwloder

import (
	"OMPFinex-CodeChallenge/internal/contract/image"
	"OMPFinex-CodeChallenge/pkg/log"
	"OMPFinex-CodeChallenge/services/manager/interactor/collector"
)

type Manager struct {
	imageRepo       image.Repository
	logger          log.Logger
	imageChannel    chan string
	imageCollectors map[string]*collector.ImageCollector
}

func New(imageRepo image.Repository, imageChannel chan string, logger log.Logger) *Manager {
	m := make(map[string]*collector.ImageCollector)
	return &Manager{
		imageRepo:       imageRepo,
		logger:          logger,
		imageChannel:    imageChannel,
		imageCollectors: m,
	}
}
