package models

type UploadResponse struct {
	Message     string
	FailedCards []*CardRequest
}
