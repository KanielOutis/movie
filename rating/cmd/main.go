package main

import (
	"log"
	"net/http"

	"movieexample.com/rating/internal/controller/rating"
	httphandler "movieexample.com/rating/internal/handler/http"
	"movieexample.com/rating/internal/repository/memory"
)

func main() {
	log.Printf("Starting rating service")
	repository := memory.New()
	controller := rating.New(repository)
	handler := httphandler.New(controller)
	http.Handle("/rating", http.HandlerFunc(handler.Handle))

	if err := http.ListenAndServe(":3001", nil); err != nil {
		panic(err)
	}
}
