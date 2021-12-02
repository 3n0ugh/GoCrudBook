package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/3n0ugh/GoCrudBook/cmd/web/config"
	"github.com/3n0ugh/GoCrudBook/pkg/models"
)

// home
func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

// Get all books
func BookGetAll(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		books, err := app.Books.GetAll()
		if err != nil {
			if errors.Is(err, models.ErrNoRecord) {
				app.NotFound(w)
				return
			}
			app.ServerError(w, err)
			return
		}

		for _, b := range books {
			fmt.Fprintf(w, "%v\n\n", b)
		}
	}
}

// Get book by id
func BookGetById(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil || id < 1 {
			app.NotFound(w)
		}

		book, err := app.Books.GetById(id)
		if err != nil {
			if errors.Is(err, models.ErrNoRecord) {
				app.NotFound(w)
				return
			}
			app.ServerError(w, err)
			return
		}

		fmt.Fprintf(w, "%v\n", book)
	}
}

// Get book by name
func BookGetByName(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

// Add book
func BookAdd(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

// Delete book
func BookDelete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

// Update book
func BookUpdate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

// Borrow book
func BookBorrow(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

// Receive book
func BookReceive(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}
