package models

import (
	"errors"

	"github.com/go-sql-driver/mysql"
)

var ErrNoRecord = errors.New("models: no record found")

type Book struct {
	ISBN            int            `json:"isbn"`
	BookName        string         `json:"book_name"`
	Author          string         `json:"author"`
	PageCount       int            `json:"page_count"`
	BookCount       int            `json:"book_count"`
	BorrowTimes     int            `json:"borrow_times"`
	BorrowDate      mysql.NullTime `json:"borrow_date"`
	LastReceivedDay mysql.NullTime `json:"last_received_day"`
}
