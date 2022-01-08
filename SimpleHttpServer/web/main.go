package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func handlerHello(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "text/plain")
	rw.Write([]byte("Hello World!"))
}

func main() {
	port := os.Getenv("PORT")

	log.Printf("Starting a go-117 server")

	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	r := chi.NewRouter()
	r.Get("/hello", handlerHello)
	log.Printf("Listening on port %s", port)

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}

	log.Printf("Stopping a go-117 server")
}
