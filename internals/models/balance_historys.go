package models

import "time"

type BalanceHistory struct {
	Id         int64
	User       int64
	Type       string
	Amount     float64
	Before     float64
	After      float64
	References float64
	CreatedAt  time.Time
}

type CreateBalanceHistory struct {
	User       int64
	Type       string
	Amount     float64
	Before     float64
	After      float64
	References float64
}
