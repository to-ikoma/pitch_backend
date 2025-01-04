package model

import "time"

type Date struct {
	Year  int32
	Month int32
	Day   int32
}

func NewDate() *Date {
	now := time.Now()

	return &Date{
		Year:  int32(now.Year()),
		Month: int32(now.Month()),
		Day:   int32(now.Day()),
	}
}
