package models

type DownloadRequest struct {
	Manga   string `json:"manga" validate:"required,oneof=one-piece vinland-saga"`
	Chapter int    `json:"chapter" validate:"required,min=1,max=2000"`
}
