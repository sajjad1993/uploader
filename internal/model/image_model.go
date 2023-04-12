package model

import "time"

type Image struct {
	Sha256    string    `json:"sha"`
	Size      int       `json:"size"`
	ChunkSize int       `json:"chunk_size"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type ImageStatus string

const (
	UnCompletedStatus ImageStatus = "uncompleted"
	CompletedStatus   ImageStatus = "completed"
	ReadyStatus       ImageStatus = "ready"
)
