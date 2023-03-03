package model

import "time"

type Item struct {
	Title       string    `json:"title"`
	Source      string    `json:"source"`
	SourceUrl   string    `json:"source_url"`
	Link        string    `json:"link"`
	PublishDate time.Time `json:"publish_date"`
	Description string    `json:"description"`
}
