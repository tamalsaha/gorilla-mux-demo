package main

import (
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "XYZ")
	return
}

func main() {
	r := mux.NewRouter()
	r.Host("")
	r.HandleFunc("/", handler)
	r.HandleFunc("/products", handler).Methods("POST")
	r.HandleFunc("/articles", handler).Methods("GET")
	r.HandleFunc("/articles/{id}", handler).Methods("GET", "PUT")
	r.HandleFunc("/authors", handler).Queries("surname", "{surname}")

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
