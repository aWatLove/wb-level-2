package http

import (
	"dev_11/internal/model"
	"log"
	"net/http"
	"time"
)

func (h *Handler) createEvent(w http.ResponseWriter, r *http.Request) {
	if httpMethodErrorCheck(w, http.MethodPost, r.Method) {
		return
	}

	input, err := h.decodeCreateEventBodyJSON(r)
	if err != nil {
		h.httpErrorResponse(w, http.StatusBadRequest, err.Error())
		log.Println(err.Error())
		return
	}

	err = h.service.CreateEvent(input.UserID, model.Event{
		Name:        input.Name,
		Description: input.Description,
		Date:        time.Time(input.Date),
	})
	if err != nil {
		h.httpErrorResponse(w, http.StatusServiceUnavailable, err.Error())
		log.Println(err.Error())
		return
	}

	h.httpSuccessEventActionResponse(w, http.StatusOK, "Event was created")
}

func (h *Handler) updateEvent(w http.ResponseWriter, r *http.Request) {
	if httpMethodErrorCheck(w, http.MethodPost, r.Method) {
		return
	}

	input, err := h.decodeUpdateEventBodyJSON(r)
	if err != nil {
		h.httpErrorResponse(w, http.StatusBadRequest, err.Error())
		log.Println(err.Error())
		return
	}

	err = h.service.UpdateEvent(input.UserID, model.Event{
		ID:          input.EventID,
		Name:        input.Name,
		Description: input.Description,
		Date:        time.Time(input.Date),
	})
	if err != nil {
		h.httpErrorResponse(w, http.StatusServiceUnavailable, err.Error())
		log.Println(err.Error())
		return
	}

	h.httpSuccessEventActionResponse(w, http.StatusOK, "Event was updated")
}

func (h *Handler) deleteEvent(w http.ResponseWriter, r *http.Request) {
	if httpMethodErrorCheck(w, http.MethodPost, r.Method) {
		return
	}

	input, err := h.decodeDeleteEventBodyJSON(r)
	if err != nil {
		h.httpErrorResponse(w, http.StatusBadRequest, err.Error())
		log.Println(err.Error())
		return
	}

	err = h.service.DeleteEvent(input.UserID, input.EventID)
	if err != nil {
		h.httpErrorResponse(w, http.StatusServiceUnavailable, err.Error())
		log.Println(err.Error())
		return
	}

	h.httpSuccessEventActionResponse(w, http.StatusOK, "Event was deleted")
}

func (h *Handler) getEventsForDay(w http.ResponseWriter, r *http.Request) {
	if httpMethodErrorCheck(w, http.MethodGet, r.Method) {
		return
	}

	userID, date, err := getParamsInput(r.URL)
	if err != nil {
		h.httpErrorResponse(w, http.StatusBadRequest, err.Error())
		log.Println(err.Error())
		return
	}

	events, err := h.service.GetEventsForDay(userID, date)
	if err != nil {
		h.httpErrorResponse(w, http.StatusServiceUnavailable, err.Error())
		log.Println(err.Error())
		return
	}

	h.httpEventsResponse(w, http.StatusOK, events)
}

func (h *Handler) getEventsForWeek(w http.ResponseWriter, r *http.Request) {
	if httpMethodErrorCheck(w, http.MethodGet, r.Method) {
		return
	}

	userID, date, err := getParamsInput(r.URL)
	if err != nil {
		h.httpErrorResponse(w, http.StatusBadRequest, err.Error())
		log.Println(err.Error())
		return
	}

	events, err := h.service.GetEventsForWeek(userID, date)
	if err != nil {
		h.httpErrorResponse(w, http.StatusServiceUnavailable, err.Error())
		log.Println(err.Error())
		return
	}

	h.httpEventsResponse(w, http.StatusOK, events)
}

func (h *Handler) getEventsForMonth(w http.ResponseWriter, r *http.Request) {
	if httpMethodErrorCheck(w, http.MethodGet, r.Method) {
		return
	}

	userID, date, err := getParamsInput(r.URL)
	if err != nil {
		h.httpErrorResponse(w, http.StatusBadRequest, err.Error())
		log.Println(err.Error())
		return
	}

	events, err := h.service.GetEventsForMonth(userID, date)
	if err != nil {
		h.httpErrorResponse(w, http.StatusServiceUnavailable, err.Error())
		log.Println(err.Error())
		return
	}

	h.httpEventsResponse(w, http.StatusOK, events)
}
