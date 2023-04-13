package collector

import (
	"OMPFinex-CodeChallenge/internal/contract/chunk"
	"OMPFinex-CodeChallenge/internal/contract/image"
	"OMPFinex-CodeChallenge/pkg/log"
	rpc "OMPFinex-CodeChallenge/pkg/rpc/proto"
	"OMPFinex-CodeChallenge/services/manager/entity"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type Collector struct {
	//todo proto adapter
	ImageRepo       image.Reader
	chunkRepo       chunk.Repository
	imageChannel    chan string
	logger          log.Logger
	imageCollectors map[string]*ImageCollector
	client          rpc.MergerClient
	timeout         time.Duration
}

// ImageCollector  get an image and gather its  chunks
type ImageCollector struct {
	chunkRepo             chunk.Repository
	logger                log.Logger
	chunkChannel          chan entity.Chunk
	completedImageChannel chan string
	sha256                string
	completedChunk        int
	ChunkCount            int
	timeout               time.Duration
}

func NewImageCollector(image entity.Image, completedImageChannel chan string, chunkRepo chunk.Repository, logger log.Logger, timeout time.Duration) *ImageCollector {
	chunkCount := getChunkCount(image)
	chunkChannel := make(chan entity.Chunk, chunkCount)
	return &ImageCollector{
		sha256:                image.Sha256,
		completedChunk:        0,
		ChunkCount:            chunkCount,
		chunkChannel:          chunkChannel,
		completedImageChannel: completedImageChannel,
		chunkRepo:             chunkRepo,
		logger:                logger,
		timeout:               timeout,
	}

}
func New(imageRepo image.Repository, chunkRepo chunk.Repository, MergerAddress string, logger log.Logger, timeout time.Duration) (UseCase, error) {
	channel := make(chan string)
	imageCollectors := make(map[string]*ImageCollector)
	cc, err := grpc.Dial(fmt.Sprintf("%s", MergerAddress), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatal(fmt.Sprintf("could not connect to server so it will be a panic: %v", err))
		cc.Close()

		return nil, err
	}
	client := rpc.NewMergerClient(cc)
	return &Collector{
		ImageRepo:       imageRepo,
		logger:          logger,
		imageChannel:    channel,
		imageCollectors: imageCollectors,
		chunkRepo:       chunkRepo,
		client:          client,
		timeout:         timeout,
	}, nil
}
