package service

import (
	"dev_11/internal/model"
	"dev_11/internal/repository"
	"time"
)

// User - интерфейс для сервиса пользователя
type User interface {
	GetEventsForDay(id string, date time.Time) ([]model.Event, error)
	GetEventsForWeek(id string, startWeekDate time.Time) ([]model.Event, error)
	GetEventsForMonth(id string, startMonthDate time.Time) ([]model.Event, error)
	CreateEvent(userID string, event model.Event) error
	UpdateEvent(userID string, event model.Event) error
	DeleteEvent(userID, eventID string) error
}

// Service - структура реализующая интерфейс User
type Service struct {
	cache repository.UserCacheRepo
}

// NewService - конструктор Service
func NewService() *Service {
	return &Service{cache: *repository.NewUserCacheRepo(repository.NewCache())}
}
