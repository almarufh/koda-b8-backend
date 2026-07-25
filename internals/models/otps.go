package models

import "time"

type Otp struct {
	Id        int64
	User      int64
	Code      string
	Used      bool
	CreatedAt time.Time
	ExpiredAt time.Time
}

type GenerateOtp struct {
	User int64
	Code string
}
