package model

type ShortenURL struct {
	ID      uint64 `json:"url_id"`
	Token   string `json:"token"`
	FullURL string `json:"full_url"`
}
