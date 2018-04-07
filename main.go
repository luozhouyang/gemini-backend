package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	config := NewConfig()
	files := http.FileServer(http.Dir(config.Static))
	mux.Handle("/static/", http.StripPrefix("/static", files))

	server := &http.Server{
		Addr:    config.Address,
		Handler: mux,
	}
	server.ListenAndServe()
}
