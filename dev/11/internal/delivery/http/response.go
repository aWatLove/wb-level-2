package http

import (
	"dev_11/internal/model"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) httpSuccessEventActionResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	data := newSuccessEventOutput(message)
	response, _ := json.MarshalIndent(data, " ", "")

	_, err := w.Write(response)
	if err != nil {
		http.Error(w, fmt.Errorf("error: %v", err).Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) httpEventsResponse(w http.ResponseWriter, statusCode int, events []model.Event) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	data := newEventsOutput(events)
	response, _ := json.MarshalIndent(data, " ", "")

	_, err := w.Write(response)
	if err != nil {
		http.Error(w, fmt.Errorf("error: %v", err).Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) httpErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	data := newErrorOutput(message)
	response, _ := json.MarshalIndent(data, " ", "")

	_, err := w.Write(response)
	if err != nil {
		http.Error(w, fmt.Errorf("error: %v", err).Error(), http.StatusInternalServerError)
	}
}
