package models

import (
	"errors"

	"github.com/go-sql-driver/mysql"
)

var ErrNoRecord = errors.New("models: no record found")

type Book struct {
	ISBN            int
	BookName        string
	Author          string
	PageCount       int
	BookCount       int
	BorrowTimes     int
	BorrowDate      mysql.NullTime
	LastReceivedDay mysql.NullTime
}
