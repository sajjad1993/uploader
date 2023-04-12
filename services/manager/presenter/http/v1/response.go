package v1

// Response represents the structure that our APIs respond in that way
type Response struct {
	ErrorDetail map[string]string `json:"error-detail"`
	Error       string            `json:"error"`
	Message     string            `json:"message"`
	Data        interface{}       `json:"data"`
}
