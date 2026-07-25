package models

import "time"

type User struct {
	Id        int64
	Email     string
	Password  string
	Pin       string
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
	CreatedBy time.Time
}

type Register struct {
	Email    string
	Password string
	Pin      string
	Phone    string
}
