package mysql

import (
	"database/sql"
	"errors"

	"github.com/3n0ugh/GoCrudBook/cmd/pkg/models"
)

type BookModel struct {
	DB *sql.DB
}

func (b *BookModel) GetAll() ([]*models.Book, error) {
	stmt := `SELECT * FROM book`
	rows, err := b.DB.Query(stmt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		}
		return nil, err
	}

	books := []*models.Book{}

	for rows.Next() {
		b := &models.Book{}
		err = rows.Scan(
			&b.ISBN, &b.BookName, &b.Author, &b.PageCount,
			&b.BookCount, &b.BorrowTimes, &b.BorrowDate, &b.LastReceivedDay)
		if err != nil {
			return nil, err
		}

		books = append(books, b)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func (b *BookModel) GetById(id int) (*models.Book, error) {
	stmt := `SELECT * FROM book WHERE ISBN = ?`

	row := b.DB.QueryRow(stmt, id)

	book := &models.Book{}
	err := row.Scan(
		&book.ISBN, &book.BookName, &book.Author, &book.PageCount,
		&book.BookCount, &book.BorrowTimes, &book.BorrowDate, &book.LastReceivedDay)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			return nil, models.ErrNoRecord
		}
		return nil, err
	}

	return book, nil
}
