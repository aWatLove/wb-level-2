package service

import (
	"dev_11/internal/model"
	"dev_11/internal/repository"
	"time"
)

type User interface {
	GetEventsForDay(id string, date time.Time) ([]model.Event, error)
	GetEventsForWeek(id string, startWeekDate time.Time) ([]model.Event, error)
	GetEventsForMonth(id string, startMonthDate time.Time) ([]model.Event, error)
	CreateEvent(userId string, event model.Event) error
	UpdateEvent(userId string, event model.Event) error
	DeleteEvent(userId, eventId string) error
}

type Service struct {
	cache repository.UserCacheRepo
}

func NewService() *Service {
	return &Service{cache: *repository.NewUserCacheRepo(repository.NewCache())}
}
