package entity

import (
	"OMPFinex-CodeChallenge/internal/model"
)

type Image struct {
	Sha256    string `json:"sha256"`
	Size      int    `json:"size"`
	ChunkSize int    `json:"chunk_size"`
}
type Chunk struct {
	ID     int    `json:"id"`
	Size   int    `json:"size"`
	Data   string `json:"data"`
	Sha256 string `json:"sha256"`
}

func ImageEntityToModel(image Image) model.Image {
	return model.Image{
		Sha256:    image.Sha256,
		Size:      image.Size,
		ChunkSize: image.ChunkSize,
	}
}
func ImageModelToEntity(image model.Image) Image {
	return Image{
		Sha256:    image.Sha256,
		Size:      image.Size,
		ChunkSize: image.ChunkSize,
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
