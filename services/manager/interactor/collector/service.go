package collector

import (
	"OMPFinex-CodeChallenge/pkg/errs"
	rpc "OMPFinex-CodeChallenge/pkg/rpc/proto"
	"OMPFinex-CodeChallenge/services/manager/entity"
	"context"
	"fmt"
	"math"
	"os"
)

// StartCollecting create new image collector
func (c *Collector) StartCollecting(ctx context.Context, image entity.Image) {
	// create image collector
	imageCollector := NewImageCollector(image, c.imageChannel, c.chunkRepo, c.logger, c.timeout)
	c.imageCollectors[image.Sha256] = imageCollector
	go imageCollector.Gather()
}

// SaveChunk gathers chunks until they get ready for merging
func (c *Collector) SaveChunk(ctx context.Context, chunk entity.Chunk) error {
	ok, err := c.ImageRepo.DoesExist(ctx, chunk.Sha256)
	if err != nil {
		return err
	}
	if !ok {
		return errs.NewValidationError("The image dose not exist")
	}
	imageCollector := c.imageCollectors[chunk.Sha256]
	if imageCollector == nil {
		return errs.NewValidationError("The image dose not exist")
	}
	imageCollector.chunkChannel <- chunk
	return nil
}

func (c *Collector) CallMerger() {
	for imageHash := range c.imageChannel {
		//todo call grpc
		ctx, _ := context.WithTimeout(context.Background(), c.timeout)
		c.logger.Info(fmt.Sprintf("image %s has been completed", imageHash))
		_, err := c.client.Merge(ctx, &rpc.MergeRequest{Image: &rpc.Image{Sha256: imageHash}})
		if err != nil {
			c.logger.Error(fmt.Sprintf("image %s has been completed", imageHash))
		}

	}
}

func (c *Collector) GetChannel() chan string {
	return c.imageChannel
}
func (ic *ImageCollector) Gather() {
	fmt.Println("upload input ")
	for chunkEnt := range ic.chunkChannel {

		err := ic.writeChunk(chunkEnt)
		if err != nil {
			//todo handel error
			ic.logger.Error(err.Error())
		}
		chunkEnt.Data = chunkEnt.FileAddress()
		err = ic.saveChunk(chunkEnt)
		if err != nil {
			ic.logger.Error(err.Error())
		}
		ic.completedChunk++
		if ic.completedChunk == ic.ChunkCount {
			ic.completedImageChannel <- ic.sha256
		}

	}
}

func getChunkCount(image entity.Image) int {
	a, b := float64(image.Size), float64(image.ChunkSize)
	if math.Mod(a, b) == 0 {
		return int(a / b)
	}
	return int(math.Abs(a/b)) + 1
}

func (ic *ImageCollector) writeChunk(chunk entity.Chunk) error {
	b := []byte(chunk.Data)
	fmt.Printf("%s", b)
	err := os.MkdirAll(chunk.FileDir(), 0777)

	if err != nil && !os.IsExist(err) {
		return err
	}
	//filepath.Join(filepath.F)
	err = os.WriteFile(chunk.FileAddress(), b, os.ModePerm)
	return err

}
func (ic *ImageCollector) saveChunk(chunk entity.Chunk) error {
	ctx, cancel := context.WithTimeout(context.Background(), ic.timeout)
	defer cancel()
	chunkModel := entity.ChunkEntityToModel(chunk)
	err := ic.chunkRepo.Save(ctx, chunkModel)
	if err != nil {
		return err
	}
	return nil
}
