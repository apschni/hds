package dto

import "mime/multipart"

type UploadFileOnTaskReq struct {
	Id   string                `uri:"id" binding:"required"`
	File *multipart.FileHeader `form:"file" binding:"required"`
}
