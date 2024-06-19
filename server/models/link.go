package models

type Link struct {
	ID        int    `json:"id"`
	URL       string `json:"url"`
	Platform  string `json:"platform"`
	OderIndex int    `json:"order_index"`
}
