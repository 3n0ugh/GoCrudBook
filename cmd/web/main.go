package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	port := flag.String("port", ":5000", "Port Address")
	flag.Parse()

	srv := &http.Server{
		Addr: *port,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
