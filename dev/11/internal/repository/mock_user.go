package repository

import "dev_11/internal/model"

func (o *UserCacheRepo) addTestUser() {
	testUser := model.NewUser("1")
	o.PutUser(testUser.Id, testUser)
}
