package model

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Category struct {
	Id       uuid.UUID      `json:"id" db:"id"`
	Label    string         `json:"label" db:"label"`
	Subjects pq.StringArray `json:"subjects" db:"subjects"`
}
