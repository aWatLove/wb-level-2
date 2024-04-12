package repository

import (
	"dev_11/internal/model"
	"sync"
)

type Cache struct {
	Mutex sync.RWMutex
	Data  map[string]model.User
}

func NewCache() *Cache {
	return &Cache{Data: make(map[string]model.User)}
}
