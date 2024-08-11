package main

import (
	"log"
	"net/http"

	"github.com/amcollie/htmx-for-beginners-go/book"
	"github.com/amcollie/htmx-for-beginners-go/middleware"
)

func main() {
	port := ":8080"

	bookHandler := &book.Handler{}

	mux := http.NewServeMux()

	// mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.HandleFunc("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))).ServeHTTP)
	mux.HandleFunc("GET /", bookHandler.Index)
	mux.HandleFunc("GET /books/edit/{id}", bookHandler.Edit)
	mux.HandleFunc("GET /books", bookHandler.Show)
	mux.HandleFunc("POST /books", bookHandler.Store)
	mux.HandleFunc("GET /books/{id}", bookHandler.Details)
	mux.HandleFunc("PUT /books/{id}", bookHandler.Update)
	mux.HandleFunc("DELETE /books/{id}", bookHandler.Destroy)
	mux.HandleFunc("POST /books/search", bookHandler.Find)

	middlewareStack := middleware.CreateStack(
		middleware.Logging,
	)

	server := http.Server{
		Addr:    port,
		Handler: middlewareStack(mux),
	}

	log.Printf("Listening on %s", port)

	server.ListenAndServe()
}
