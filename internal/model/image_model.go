package model

import "time"

type Image struct {
	Sha256    string    `json:"sha"`
	Size      uint      `json:"size"`
	ChunkSize uint      `json:"chunk_size"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	Data      string    `json:"data"`
}

type ImageStatus string

const (
	UnCompletedStatus ImageStatus = "uncompleted"
	CompletedStatus   ImageStatus = "completed"
	ReadyStatus       ImageStatus = "ready"
)
