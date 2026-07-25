package models

import (
	"time"
)

type Profile struct {
	Id        int64
	IdUser    int64
	FirstName string
	LastName  string
	Birthday  time.Time
	Mother    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
