package router

import (
	"net/http"

	"github.com/3n0ugh/GoCrudBook/cmd/web/handler"
)

func SetRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.Home)

	return mux
}
