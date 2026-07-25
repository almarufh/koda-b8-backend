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

type ChangeEmail struct {
	Email   string
	Confirm string
	otp     string
}

type ChangePassword struct {
	Password string
	Confirm  string
	otp      string
}

type ChangePin struct {
	Pin     string
	Confirm string
	otp     string
}

type ChangePhone struct {
	Phone   string
	Confirm string
	otp     string
}
