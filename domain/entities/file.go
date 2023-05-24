package entities

type UploadResponse struct {
	Message     string
	FailedCards []*CardRequest
}
