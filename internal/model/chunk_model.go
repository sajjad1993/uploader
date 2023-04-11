package model

type Chunk struct {
	ID     int    `json:"id"`
	Size   int    `json:"size"`
	Data   string `json:"data"`
	Sha256 string `json:"sha256"`
}
