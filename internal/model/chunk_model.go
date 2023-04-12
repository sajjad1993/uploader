package model

import "time"

type Chunk struct {
	ID        int       `json:"id"`
	Size      int       `json:"size"`
	Data      string    `json:"data"`
	Sha256    string    `json:"sha"`
	CreatedAt time.Time `json:"created_at"`
}
