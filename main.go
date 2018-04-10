package main

import (
	"net/http"
	"backend/handlers"
)

func main() {
	mux := http.NewServeMux()
	config := NewConfig()
	files := http.FileServer(http.Dir(config.Static))
	mux.Handle("/static/", http.StripPrefix("/static", files))

	mux.HandleFunc("/login", handlers.Login)
	mux.HandleFunc("/signup", handlers.SignUp)

	server := &http.Server{
		Addr:    config.Address,
		Handler: mux,
	}
	server.ListenAndServe()
}
