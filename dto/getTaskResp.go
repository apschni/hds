package dto

import (
	"github.com/google/uuid"
	"time"
)

type GetTaskResp struct {
	Id         uuid.UUID `json:"id"`
	CategoryId uuid.UUID `json:"category_id" db:"category_id"`
	SubjectId  uuid.UUID `json:"subject_id" db:"subject_id"`
	Label      string    `json:"label"`
	Text       string    `json:"text"`
	Subject    string    `json:"subject"`
	Keypoint   bool      `json:"keypoint"`
	Points     int       `json:"points"`
	Completed  bool      `json:"completed"`
	Deadline   time.Time `json:"deadline"`
}
