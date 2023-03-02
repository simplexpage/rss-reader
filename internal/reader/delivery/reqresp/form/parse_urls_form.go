package form

type ParseUrlsForm struct {
	Urls []string `json:"urls" validate:"required"`
}
