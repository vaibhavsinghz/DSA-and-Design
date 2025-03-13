package service

import (
	"Self/splitwise/models"
	"sync"
)

type UserService struct {
	Users      map[int]*models.User
	nextUserId int
}

var (
	userService *UserService
	onceUser    sync.Once
)

func NewUserService() *UserService {
	onceUser.Do(func() {
		userService = &UserService{
			Users:      make(map[int]*models.User),
			nextUserId: 1,
		}
	})
	return userService
}

func (u *UserService) RegisterUser(name string) int {
	user := models.NewUser(u.nextUserId, name)
	u.Users[user.ID] = user
	u.nextUserId++
	return user.ID
}

func (u *UserService) DeleteUser(userId int) {
	delete(u.Users, userId)
}

func (u *UserService) GetUserById(userId int) *models.User {
	return u.Users[userId]
}
