package main

import (
	"flag"
	"log"
	"net/http"
)

type application struct{}

func main() {
	addr := flag.String("addr", ":9000", "HTTP Network")

	flag.Parse()

	app := &application{}

	srv := http.Server{
		Addr:    *addr,
		Handler: app.Routes(),
	}

	log.Printf("Server starting on %s", *addr)
	err := srv.ListenAndServe()

	log.Fatal(err)
}
