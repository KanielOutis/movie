package main

import (
	"log"
	"net/http"

	"movieexample.com/metadata/internal/controller/metadata"
	httphandler "movieexample.com/metadata/internal/handler"
	"movieexample.com/metadata/internal/repository/memory"
)

func main() {
	log.Printf("Starting metadata service")

	repo := memory.New()
	controller := metadata.New(repo)
	h := httphandler.New(controller)

	http.Handle("/metadata", http.HandlerFunc(h.GetMetadata))

	if err := http.ListenAndServe(":3000", nil); err != nil {
		panic(err)
	}
}
