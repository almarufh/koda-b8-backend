package models

import "time"

type Balance struct {
	User      int64
	Balance   float64
	UpdatedAt time.Time
}
