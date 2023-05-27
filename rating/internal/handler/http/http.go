package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"movieexample.com/rating/internal/controller/rating"
	"movieexample.com/rating/pkg/model"
)

type Handler struct {
	controller *rating.Controller
}

func New(controller *rating.Controller) *Handler {
	return &Handler{
		controller: controller,
	}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	recordID := model.RecordID(r.FormValue("record_id"))
	recordType := model.RecordType(r.FormValue("record_type"))

	if recordID == "" || recordType == "" {
		http.Error(w, "recordID and recordType are required", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		v, err := h.controller.GetAggregatedRating(r.Context(), model.RecordID(r.FormValue("record_id")), model.RecordType(r.FormValue("record_type")))
		if err != nil && errors.Is(err, rating.ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		} else if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if err := json.NewEncoder(w).Encode(v); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case http.MethodPut:
		v, err := strconv.ParseFloat(r.FormValue("value"), 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err := h.controller.PutRating(r.Context(), recordID, recordType, &model.Rating{
			RecordID:   recordID,
			RecordType: recordType,
			Value:      model.RatingValue(v),
		}); err != nil {
			log.Printf("Repository put error: %v\n", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
