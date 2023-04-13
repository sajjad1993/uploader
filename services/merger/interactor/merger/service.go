package merger

import (
	"OMPFinex-CodeChallenge/services/merger/entity"
	"bufio"
	"context"
	"io"
	"os"
	"time"
)

// MergeChunks create image based on chunks
func (m Merger) MergeChunks(ctx context.Context, sha string) error {
	imageModel, err := m.imageRepo.Get(ctx, sha)
	if err != nil {
		return err
	}
	all, err := m.chunkRepo.GetAll(ctx, sha)
	if err != nil {
		return err
	}
	var chunks []entity.Chunk
	for _, chunkEnt := range all {
		chunks = append(chunks, entity.ChunkModelToEntity(chunkEnt))
	}
	go m.mergePic(entity.ImageModelToEntity(*imageModel), chunks)
	return nil
}

func (m Merger) UpdateImage(ctx context.Context, image entity.Image) error {
	image.Data = image.FileAddress()
	image.Status = string(entity.ReadyStatus)
	return m.imageRepo.Update(ctx, entity.ImageEntityToModel(image))
}

func (m Merger) mergePic(image entity.Image, chunks []entity.Chunk) {
	err := m.prepareImage(image, chunks)
	if err != nil {
		m.logger.Error(err.Error())
		return
	}
	//todo config this
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = m.UpdateImage(ctx, image)
	if err != nil {
		m.logger.Error(err.Error())
		return
	}
}

func (m Merger) prepareImage(image entity.Image, chunks []entity.Chunk) error {
	imageFile, err := m.createImageFile(image)
	if err != nil {
		m.logger.Error(err.Error())
	}
	for _, chunkEnt := range chunks {

		chunkFile, err := os.Open(chunkEnt.Data)
		if err != nil {
			m.logger.Error(err.Error())
		}
		// make a read buffer
		r := bufio.NewReader(chunkFile)

		// make a write buffer
		w := bufio.NewWriter(imageFile)

		// make a buffer to keep chunks that are read
		buf := make([]byte, 1024)
		for {
			// read a chunk
			n, err := r.Read(buf)
			if err != nil && err != io.EOF {
				m.logger.Error(err.Error())
			}
			if n == 0 {
				break
			}

			// write a chunk
			if _, err := w.Write(buf[:n]); err != nil {
				m.logger.Error(err.Error())
			}
		}

		if err = w.Flush(); err != nil {
			m.logger.Error(err.Error())
		}
	}
	return nil
}

func (m Merger) createImageFile(image entity.Image) (*os.File, error) {
	err := os.MkdirAll(image.FileDir(), 0777)

	if err != nil && !os.IsExist(err) {
		return nil, err
	}
	//filepath.Join(filepath.F)
	file, err := os.Create(image.FileAddress())
	if err != nil && !os.IsExist(err) {
		return nil, err
	}
	return file, nil
}
