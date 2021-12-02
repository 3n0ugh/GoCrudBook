package router

import (
	"net/http"

	"github.com/3n0ugh/GoCrudBook/cmd/web/config"
	"github.com/3n0ugh/GoCrudBook/cmd/web/handler"
)

func SetRoutes(app *config.Application) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/book", handler.BookGetById(app))
	mux.HandleFunc("/book/all", handler.BookGetAll(app))

	return mux
}
