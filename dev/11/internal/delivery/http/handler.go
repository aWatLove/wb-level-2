package http

import (
	"dev_11/internal/service"
	"net/http"
)

type Handler struct {
	service service.User
}

func NewHandler(service service.User) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/create_event", h.createEvent)
	mux.HandleFunc("/update_event", h.updateEvent)
	mux.HandleFunc("/delete_event", h.deleteEvent)
	mux.HandleFunc("/event_for_day", h.getEventsForDay)
	mux.HandleFunc("/event_for_week", h.getEventsForWeek)
	mux.HandleFunc("/event_for_month", h.getEventsForMonth)
	handler := Log(mux)
	return handler
}
