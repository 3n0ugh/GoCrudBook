package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/3n0ugh/GoCrudBook/cmd/pkg/models/mysql"
	"github.com/3n0ugh/GoCrudBook/cmd/web/config"
	"github.com/3n0ugh/GoCrudBook/cmd/web/database"
	"github.com/3n0ugh/GoCrudBook/cmd/web/router"
)

func main() {
	port := flag.String("port", ":5000", "Port Address")
	dsn := flag.String("dsn", "webook:pass1234!@/library", "MySQL database data source name")
	flag.Parse()

	db, err := database.DBOpen(*dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &config.Application{
		InfoLog:  infoLog,
		ErrorLog: errorLog,
		Books:    &mysql.BookModel{DB: db},
	}

	srv := &http.Server{
		Addr:    *port,
		Handler: router.SetRoutes(app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
