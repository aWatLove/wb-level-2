package repository

import (
	"dev_11/internal/model"
	"sync"
)

// Cache - структура хранилища в виде кэша
type Cache struct {
	Mutex sync.RWMutex
	Data  map[string]model.User
}

// NewCache - конструктор структуры Cache
func NewCache() *Cache {
	return &Cache{Data: make(map[string]model.User)}
}
