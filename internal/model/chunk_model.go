package model

import "time"

type Chunk struct {
	ID        uint      `json:"id"`
	Size      uint      `json:"size"`
	Data      string    `json:"data"`
	Sha256    string    `json:"sha"`
	CreatedAt time.Time `json:"created_at"`
}
