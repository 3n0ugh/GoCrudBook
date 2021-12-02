package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no record found")

type Book struct {
	ISBN            int
	BookName        string
	Author          string
	PageCount       int
	BookCount       int
	BorrowTimes     int
	BorrowDate      time.Time
	LastReceivedDay time.Time
}
