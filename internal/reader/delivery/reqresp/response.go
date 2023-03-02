package reqresp

import "github.com/gookit/validate"

type ParseUrlsResponse struct {
	Items         []Item          `json:"items,omitempty"`
	Err           error           `json:"err,omitempty"`
	ValidationErr validate.Errors `json:"validation_err,omitempty"`
}

type Item struct {
	Title       string `json:"title,omitempty"`
	Source      string `json:"link,omitempty"`
	SourceUrl   string `json:"source_url,omitempty"`
	Link        string `json:"link,omitempty"`
	PublishDate string `json:"publish_date,omitempty"`
	Description string `json:"description,omitempty"`
}

func (r ParseUrlsResponse) Failed() error { return r.Err }

func (r ParseUrlsResponse) Validation() validate.Errors { return r.ValidationErr }
