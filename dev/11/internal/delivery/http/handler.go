package http

import (
	"dev_11/internal/service"
	"net/http"
)

// Handler - основной хендлер сервера
type Handler struct {
	service service.User
}

// NewHandler - конструктор Handler
func NewHandler(service service.User) *Handler {
	return &Handler{service: service}
}

// InitRoutes - метод инициализации путей
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
