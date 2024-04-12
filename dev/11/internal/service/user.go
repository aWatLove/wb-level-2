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

func (s *Service) CreateEvent(userId string, event model.Event) error {
	user, err := s.cache.GetUser(userId)
	if err != nil {
		return err
	}
	eventId := strconv.Itoa(len(user.Events) + 1)
	event.Id = eventId
	user.Events[eventId] = event
	return nil
}

func (s *Service) UpdateEvent(userId string, event model.Event) error {
	user, err := s.cache.GetUser(userId)
	if err != nil {
		return err
	}
	if _, ok := user.Events[event.Id]; !ok {
		return repository.NewErrorHandler(
			fmt.Errorf("failed to find event with id = %s of user wth id = %s", event.Id, userId),
			http.StatusBadRequest)
	}
	user.Events[event.Id] = event
	return nil
}

func (s *Service) DeleteEvent(userId, eventId string) error {
	user, err := s.cache.GetUser(userId)
	if err != nil {
		return err
	}
	if _, ok := user.Events[eventId]; !ok {
		return repository.NewErrorHandler(
			fmt.Errorf("failed to find event with id = %s of user wth id = %s", eventId, userId),
			http.StatusBadRequest)
	}
	delete(user.Events, eventId)
	return nil
}
