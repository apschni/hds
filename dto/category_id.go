package dto

import "github.com/google/uuid"

type CategoryID struct {
	ID uuid.UUID `json:"category_id" db:"category_id"`
}
