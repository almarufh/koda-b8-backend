package models

import "time"

type Otp struct {
	Id        int64
	IdUser    int64
	Code      string
	Used      bool
	CreatedAt time.Time
	ExpiredAt time.Time
}

type GenerateOtp struct {
	IdUser int64
	Code   string
}
