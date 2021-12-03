package router

import (
	"net/http"

	"github.com/3n0ugh/GoCrudBook/cmd/web/config"
	"github.com/3n0ugh/GoCrudBook/cmd/web/handler"
)

func SetRoutes(app *config.Application) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.Home(app))                   // Method Get
	mux.HandleFunc("/book", handler.BookGetAll(app))         // Method Get
	mux.HandleFunc("/book/id", handler.BookGetById(app))     // Method Get
	mux.HandleFunc("/book/name", handler.BookGetByName(app)) // Method Get
	mux.HandleFunc("/book/create", handler.BookAdd(app))     // Method Post
	mux.HandleFunc("/book/delete", handler.BookDelete(app))  // Method Delete
	mux.HandleFunc("/book/update", handler.BookUpdate(app))  // Method Put

	return mux
}
