package handler

import "net/http"

// home
func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

// Get all books
func BookGetAll(w http.ResponseWriter, r *http.Request) {

}

// Get book by id
func BookGetById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
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
