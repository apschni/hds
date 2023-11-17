package dto

import (
	"github.com/google/uuid"
	"time"
)

type CreateTaskReq struct {
	Label          string    `json:"label" binding:"required"`
	Subject        string    `json:"subject" binding:"required"`
	CategoryId     uuid.UUID `json:"category_id" db:"category_id"`
	Text           string    `json:"text"`
	Deadline       time.Time `json:"deadline" binding:"required"`
	Points         int       `json:"points"`
	IsKeyPoint     bool      `json:"is_key_point" binding:"required"`
	StudentId      uuid.UUID `json:"student_id" binding:"required"`
	VariableAnswer []string  `json:"variable_answer" db:"variable_answer"`
	Answer         string    `json:"answer" db:"answer"`
}
