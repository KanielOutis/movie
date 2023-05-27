package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"movieexample.com/metadata/internal/controller/metadata"
	"movieexample.com/metadata/internal/repository"
)

type Handler struct {
	controller *metadata.Controller
}

func New(controller *metadata.Controller) *Handler {
	return &Handler{controller}
}

func (h *Handler) GetMetadata(w http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")

	if id == "" {
		log.Printf("No id provided")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx := req.Context()
	m, err := h.controller.Get(ctx, id)

	if err != nil && errors.Is(err, repository.ErrorNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("Repository get error %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(m); err != nil {
		log.Printf("Error while json encoding %v\n", err)
	}
}
