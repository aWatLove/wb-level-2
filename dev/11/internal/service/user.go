package service

import (
	"dev_11/internal/model"
	"dev_11/internal/repository"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

const (
	hoursDay         = 24
	daysWeek         = 7
	averageDaysMonth = 30
)

// GetEventsForDay - получить события за день
func (s *Service) GetEventsForDay(id string, date time.Time) ([]model.Event, error) {
	user, err := s.cache.GetUser(id)
	if err != nil {
		return nil, err
	}

	events := make([]model.Event, 0)
	for _, v := range user.Events {
		if date.Truncate(hoursDay * time.Hour).Equal(v.Date.Truncate(hoursDay * time.Hour)) {
			events = append(events, v)
		}
	}
	return events, nil
}

// GetEventsForWeek - получить события за неделю
func (s *Service) GetEventsForWeek(id string, startWeekDate time.Time) ([]model.Event, error) {
	user, err := s.cache.GetUser(id)
	if err != nil {
		return nil, err
	}
	var difTime time.Duration

	events := make([]model.Event, 0)
	for _, v := range user.Events {
		difTime = v.Date.Sub(startWeekDate)
		if difTime > 0 && difTime < time.Hour*hoursDay*daysWeek {
			events = append(events, v)
		}
	}
	return events, nil
}

// GetEventsForMonth - получить события за месяц
func (s *Service) GetEventsForMonth(id string, startMonthDate time.Time) ([]model.Event, error) {
	user, err := s.cache.GetUser(id)
	if err != nil {
		return nil, err
	}
	var difTime time.Duration

	events := make([]model.Event, 0)
	for _, v := range user.Events {
		difTime = v.Date.Sub(startMonthDate)
		if difTime > 0 && difTime < time.Hour*hoursDay*averageDaysMonth {
			events = append(events, v)
		}
	}
	return events, nil
}

// CreateEvent - создать событие
func (s *Service) CreateEvent(userID string, event model.Event) error {
	user, err := s.cache.GetUser(userID)
	if err != nil {
		return err
	}
	eventID := strconv.Itoa(len(user.Events) + 1)
	event.ID = eventID
	user.Events[eventID] = event
	return nil
}

// UpdateEvent - обновить событие
func (s *Service) UpdateEvent(userID string, event model.Event) error {
	user, err := s.cache.GetUser(userID)
	if err != nil {
		return err
	}
	if _, ok := user.Events[event.ID]; !ok {
		return repository.NewErrorHandler(
			fmt.Errorf("failed to find event with id = %s of user wth id = %s", event.ID, userID),
			http.StatusBadRequest)
	}
	user.Events[event.ID] = event
	return nil
}

// DeleteEvent - удалить событие
func (s *Service) DeleteEvent(userID, eventID string) error {
	user, err := s.cache.GetUser(userID)
	if err != nil {
		return err
	}
	if _, ok := user.Events[eventID]; !ok {
		return repository.NewErrorHandler(
			fmt.Errorf("failed to find event with id = %s of user wth id = %s", eventID, userID),
			http.StatusBadRequest)
	}
	delete(user.Events, eventID)
	return nil
}
