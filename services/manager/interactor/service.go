package interactor

import (
	"OMPFinex-CodeChallenge/internal/contract/chunk"
	"OMPFinex-CodeChallenge/pkg/log"
	"OMPFinex-CodeChallenge/services/manager/entity"
	"context"
	"fmt"
	"time"
)

// RegisterImage register images to merges theirs chunks
func (m *Manager) RegisterImage(ctx context.Context, image entity.Image) error {
	//todo create channel
	imageModel := entity.ImageEntityToModel(image)
	err := m.imageRepo.Save(ctx, imageModel)
	if err != nil {
		m.logger.Error(err.Error())
		return err
	}
	imageCollector := newImageCollector(image, m.ImageChannel)
	go imageCollector.
	//todo check duplicate image
	//todo handle error
	// save to repo
	return nil
}

// SaveChunk gathers chunks until they get ready for merging
func (m *Manager) SaveChunk(ctx context.Context, chunk entity.Chunk) error { return nil }

// GetImage Gives image base64
func (m *Manager) GetImage() {}

func getImageChannel(image entity.Image) chan<- entity.Chunk {
	//todo get count perfect
	var chunkCount float64
	chunkCount = float64(image.Size / image.Size)
	chunkChan := make(chan entity.Chunk, chunkCount)
	return chunkChan

}

func (m *Manager) gatherChunks(chunksChannel <-chan entity.Chunk) {
	for chunk := range chunksChannel {
		//todo create config
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		err := m.SaveChunk(ctx, chunk)
		m.logger.Error(fmt.Sprintf("error %s / %s", chunk.Sha256, err))

	}
}

// ImageCollector  get an image and gather its  chunks
type ImageCollector struct {
	chunkRepo             chunk.Writer
	logger                log.Logger
	chunkChannel          chan entity.Chunk
	completedImageChannel chan entity.Image
	sha256                string
	completedChunk        int
	ChunkCount            int
}

func newImageCollector(image entity.Image, completedImageChannel chan entity.Image) *ImageCollector {
	chunkCount := getChunkSize(image)
	chunkChannel := make(chan entity.Chunk, chunkCount)
	return &ImageCollector{
		sha256:                image.Sha256,
		completedChunk:        0,
		ChunkCount:            chunkCount,
		chunkChannel:          chunkChannel,
		completedImageChannel: completedImageChannel,
	}

}

func (ic *ImageCollector) Gather() {
	for chunk := range ic.chunkChannel {
		err := ic.saveChunk(chunk)
		if err != nil {
			ic.logger.Error(err.Error())
		}
	}
}
func getChunkSize(image entity.Image) int {
	return 0
}

func (ic *ImageCollector) saveChunk(chunk entity.Chunk) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	chunkModel := entity.ChunkEntityToModel(chunk)
	err := ic.chunkRepo.Save(ctx, chunkModel)
	if err != nil {
		return err
	}
	return nil
}
