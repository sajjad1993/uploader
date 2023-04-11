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

	imageModel := entity.ImageEntityToModel(image)
	//check duplicate image
	err := m.imageRepo.CheckDuplicate(ctx, imageModel)
	if err != nil {
		m.logger.Error(err.Error())
		return err
	}
	// save to repo
	err = m.imageRepo.Save(ctx, imageModel)
	if err != nil {
		m.logger.Error(err.Error())
		return err
	}
	// create image collector
	imageCollector := newImageCollector(image, m.imageChannel)
	m.imageCollectors[image.Sha256] = imageCollector
	go imageCollector.Gather()

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
	for chunkEnt := range chunksChannel {
		//todo create config
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		err := m.SaveChunk(ctx, chunkEnt)
		m.logger.Error(fmt.Sprintf("error %s / %s", chunkEnt.Sha256, err))

	}
}

// ImageCollector  get an image and gather its  chunks
type ImageCollector struct {
	chunkRepo             chunk.Writer
	logger                log.Logger
	chunkChannel          chan entity.Chunk
	completedImageChannel chan string
	sha256                string
	completedChunk        int
	ChunkCount            int
}

func newImageCollector(image entity.Image, completedImageChannel chan string) *ImageCollector {
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
	for chunkEnt := range ic.chunkChannel {
		err := ic.saveChunk(chunkEnt)
		if err != nil {
			ic.logger.Error(err.Error())
		}
		ic.completedChunk++
		if ic.completedChunk == ic.ChunkCount {
			ic.completedImageChannel <- ic.sha256
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
