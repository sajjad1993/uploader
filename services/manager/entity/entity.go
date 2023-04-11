package entity

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
