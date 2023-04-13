package entity

import (
	"OMPFinex-CodeChallenge/internal/model"
	"OMPFinex-CodeChallenge/pkg/errs"
	"fmt"
)

type Image struct {
	Sha256    string `json:"sha256"`
	Size      uint   `json:"size"`
	ChunkSize uint   `json:"chunk_size"`
	Status    string `json:"status"`
	Data      string `json:"data"`
}
type Chunk struct {
	ID     uint   `json:"id"`
	Size   uint   `json:"size"`
	Data   string `json:"data"`
	Sha256 string `json:"sha256"`
}

func (i *Image) IsValid() error {
	if i.Size == 0 || i.ChunkSize == 0 {
		return errs.NewValidationError("size must not be zero")
	}
	return nil
}
func (c *Chunk) IsValid() error {
	if c.Size == 0 || c.Data == "" {
		return errs.NewValidationError("size must not be zero")
	}
	return nil
}

func (c *Chunk) FileDir() string {
	return fmt.Sprintf("storage/chunks/%s", c.Sha256)
}
func (c *Chunk) FileAddress() string {
	return fmt.Sprintf("%s/%d", c.FileDir(), c.ID)
}
func ImageEntityToModel(image Image) model.Image {
	return model.Image{
		Sha256:    image.Sha256,
		Size:      image.Size,
		ChunkSize: image.ChunkSize,
		Status:    image.Status,
		Data:      image.Data,
	}
}
func ImageModelToEntity(image model.Image) Image {
	return Image{
		Sha256:    image.Sha256,
		Size:      image.Size,
		ChunkSize: image.ChunkSize,
		Status:    image.Status,
		Data:      image.Data,
	}
}

func ChunkEntityToModel(chunk Chunk) model.Chunk {
	return model.Chunk{
		ID:     chunk.ID,
		Size:   chunk.Size,
		Data:   chunk.Data,
		Sha256: chunk.Sha256,
	}
}
func ChunkModelToEntity(chunk model.Chunk) Chunk {
	return Chunk{
		ID:     chunk.ID,
		Size:   chunk.Size,
		Data:   chunk.Data,
		Sha256: chunk.Sha256,
	}
}

type ImageStatus string

const (
	UnCompletedStatus ImageStatus = "uncompleted"
	CompletedStatus   ImageStatus = "completed"
	ReadyStatus       ImageStatus = "ready"
)
