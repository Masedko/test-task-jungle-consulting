package dtos

import "github.com/google/uuid"

type UploadPicture struct {
	UserID    uuid.UUID
	ImagePath string
	ImageURL  string
}

type GetImage struct {
	ImageURL string
}

type AccessImage struct {
	ImagePath string
}
