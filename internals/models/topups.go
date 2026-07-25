package models

import "time"

type Topup struct {
	Id         int64
	User       int64
	Amount     float64
	Status     string
	References string
	CreatedAt  time.Time
	ExpiredAt  time.Time
	UpdatedAt  time.Time
}

type CreateTopup struct {
	User   int64
	Amount float64
}
