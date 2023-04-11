package model

type Image struct {
	Sha256    string `json:"sha256"`
	Size      int    `json:"size"`
	ChunkSize int    `json:"chunk_size"`
	Status    string `json:"completed"`
}

type ImageStatus string

const (
	UnCompletedStatus ImageStatus = "uncompleted"
	CompletedStatus   ImageStatus = "completed"
	ReadyStatus       ImageStatus = "ready"
)
