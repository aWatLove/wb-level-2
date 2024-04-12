package repository

import (
	"dev_11/internal/model"
	"fmt"
	"net/http"
)

// UserCacheRepo - структура кеша пользователя
type UserCacheRepo struct {
	cch *Cache
}

// NewUserCacheRepo - конструктор UserCacheRepo
func NewUserCacheRepo(cch *Cache) *UserCacheRepo {
	c := UserCacheRepo{cch: cch}
	c.addTestUser()
	return &c
}

// PutUser - положить пользователя в хранилище
func (o *UserCacheRepo) PutUser(id string, user model.User) {
	o.cch.Mutex.Lock()
	defer o.cch.Mutex.Unlock()
	o.cch.Data[id] = user
}

// PutUserEvent - положить событие пользователя в хранилище
func (o *UserCacheRepo) PutUserEvent(userID string, event model.Event) error {
	o.cch.Mutex.Lock()
	defer o.cch.Mutex.Unlock()

	user, err := o.GetUser(userID)
	if err != nil {
		return err
	}

	user.Events[event.ID] = event
	return nil
}

// GetUser - получить пользователя
func (o *UserCacheRepo) GetUser(id string) (*model.User, error) {
	o.cch.Mutex.RLock()
	defer o.cch.Mutex.RUnlock()

	if userData, ok := o.cch.Data[id]; ok {
		return &userData, nil
	}
	return nil, NewErrorHandler(
		fmt.Errorf("failed to find user with id = %s", id),
		http.StatusBadRequest)
}
