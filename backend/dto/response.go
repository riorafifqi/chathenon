package dto

type Response struct {
	Success bool `json:"success"`
	Error   any  `json:"error,omitempty"`
	Data    any  `json:"data,omitempty"`
}
