package mysql

import (
	"database/sql"
	"errors"

	"github.com/3n0ugh/GoCrudBook/pkg/models"
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

func (b *BookModel) GetByName(name string) ([]*models.Book, error) {
	stmt := `SELECT * FROM book WHERE book_name LIKE ?`

	rows, err := b.DB.Query(stmt, "%"+name+"%")
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			return nil, models.ErrNoRecord
		}
		return nil, err
	}

	books := []*models.Book{}

	for rows.Next() {
		book := &models.Book{}
		err := rows.Scan(
			&book.ISBN, &book.BookName, &book.Author, &book.PageCount,
			&book.BookCount, &book.BorrowTimes, &book.BorrowDate, &book.LastReceivedDay)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func (b *BookModel) Add(book *models.Book) error {
	stmt := `INSERT INTO book VALUES(?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := b.DB.Exec(stmt,
		book.ISBN, book.BookName, book.BookName, book.PageCount, book.BookCount,
		book.BorrowTimes, book.BorrowDate, book.LastReceivedDay)
	if err != nil {
		return err
	}
	return nil
}

func (b *BookModel) Delete(id int) error {
	stmt := `DELETE FROM book WHERE isbn = ?`

	_, err := b.DB.Exec(stmt, id)
	if err != nil {
		return err
	}
	return nil
}

func (b *BookModel) Update(book *models.Book) error {
	stmt := `UPDATE book SET book_name = ?, author = ?, page_count = ?, book_count = ? ,
	borrow_times = ?, borrow_date = ?, last_received_day = ? WHERE isbn = ?`

	_, err := b.DB.Exec(stmt, book.BookName, book.Author, book.PageCount,
		book.BookCount, book.BorrowTimes, book.BorrowDate, book.LastReceivedDay, book.ISBN)
	if err != nil {
		return err
	}
	return nil
}
