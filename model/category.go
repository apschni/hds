package model

import (
	"github.com/google/uuid"
)

type Category struct {
	Id    uuid.UUID `json:"id" db:"id"`
	Label string    `json:"label" db:"label"`
}
