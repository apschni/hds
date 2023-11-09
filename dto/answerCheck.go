package dto

import (
	"github.com/google/uuid"
)

type CheckAnswer struct {
	TaskId uuid.UUID `json:"task_id" db:"task_id"`
	Answer string    `json:"answer" db:"answer"`
}
