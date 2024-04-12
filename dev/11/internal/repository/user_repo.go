package repository

import (
	"dev_11/internal/model"
	"fmt"
	"net/http"
)

type UserCacheRepo struct {
	cch *Cache
}

func NewUserCacheRepo(cch *Cache) *UserCacheRepo {
	c := UserCacheRepo{cch: cch}
	c.addTestUser()
	return &c
}

func (o *UserCacheRepo) PutUser(id string, user model.User) {
	o.cch.Mutex.Lock()
	defer o.cch.Mutex.Unlock()
	o.cch.Data[id] = user
}

func (o *UserCacheRepo) PutUserEvent(userId string, event model.Event) error {
	o.cch.Mutex.Lock()
	defer o.cch.Mutex.Unlock()

	user, err := o.GetUser(userId)
	if err != nil {
		return err
	}

	user.Events[event.Id] = event
	return nil
}

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
