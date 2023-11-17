package dto

import "github.com/lib/pq"

type GetSubjFromCategory struct {
	Subjects pq.StringArray `json:"subjects" db:"subjects"`
}
