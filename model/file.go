package model

import (
	"time"

	commonModel "github.com/ukasyah-dev/common/model"
)

type File struct {
	ID          string    `gorm:"primaryKey" json:"id"`
	Name        string    `json:"name"`
	ContentType string    `json:"contentType"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type CreateFileRequest struct {
	File string `json:"file" validate:"required"`
}

type GetFilesRequest struct {
	commonModel.PaginationRequest
}

type GetFilesResponse struct {
	commonModel.PaginationResponse
	Data []*File `json:"data"`
}

type GetFileRequest struct {
	ID string `params:"fileId" path:"fileId" validate:"required"`
}

type UpdateFileRequest struct {
	ID string `params:"fileId" path:"fileId" validate:"required"`
}

type DeleteFileRequest struct {
	ID string `params:"fileId" path:"fileId" validate:"required"`
}
