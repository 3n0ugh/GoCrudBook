package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/3n0ugh/GoCrudBook/cmd/web/router"
)

func main() {
	port := flag.String("port", ":5000", "Port Address")
	flag.Parse()

	srv := &http.Server{
		Addr:    *port,
		Handler: router.SetRoutes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
