package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world!")
	})

	port := ":"
	port = port + os.Getenv("PORT") // env var used by Heroku
	if port == ":" {
		port = ":3001"
	}

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	srv := &http.Server{
		Handler:      nil, // Default http module handler
		Addr:         port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
