package models

import (
	"time"
)

type Profile struct {
	User      int64
	FirstName string
	LastName  string
	Birthday  time.Time
	Mother    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
