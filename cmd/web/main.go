package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/3n0ugh/GoCrudBook/cmd/web/router"
)

func main() {
	port := flag.String("port", ":5000", "Port Address")
	dsn := flag.String("dsn", "webook:pass1234!@/library", "MySQL database data source name")
	flag.Parse()

	db, err := database.New(dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	srv := &http.Server{
		Addr:    *port,
		Handler: router.SetRoutes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
