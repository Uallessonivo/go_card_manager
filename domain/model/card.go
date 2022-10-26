package model

type Card struct {
	ID     string `json:"id"`
	Type   string `json:"type"`
	Owner  string `json:"owner"`
	Name   string `json:"name"`
	Serial string `json:"serial"`
}
