package models

type UrlPreview struct {
	Url         string `json:"og:url"`
	Title       string `json:"og:title"`
	Description string `json:"og:description"`
	Image       string `json:"og:image"`
	ReadingTime string `json:"og:reading_time"`
}
