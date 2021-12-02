package handler

import "net/http"

// home
func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}
