package models

import (
	"database/sql"
	"errors"
)

var ErrNoRecord = errors.New("models: no record found")

type Book struct {
	ISBN            int          `json:"isbn"`
	BookName        string       `json:"book_name"`
	Author          string       `json:"author"`
	PageCount       int          `json:"page_count"`
	BookCount       int          `json:"book_count"`
	BorrowTimes     int          `json:"borrow_times"`
	BorrowDate      sql.NullTime `json:"borrow_date"`
	LastReceivedDay sql.NullTime `json:"last_received_day"`
}
