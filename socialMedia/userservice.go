package main

import "errors"

type UserService struct {
	Users map[int]*User
}

func NewUserService() *UserService {
	return &UserService{
		Users: make(map[int]*User),
	}
}

func (us *UserService) Follow(followerID, followeeID int) error {
	follower, exist1 := us.Users[followerID]
	followee, exist2 := us.Users[followeeID]

	if !exist1 || !exist2 {
		return errors.New("user not found")
	}

	follower.Following[followeeID] = struct{}{}
	followee.Followers[followerID] = struct{}{}

	return nil
}

func (us *UserService) UnFollow(followerID, followeeID int) error {
	follower, exist1 := us.Users[followerID]
	followee, exist2 := us.Users[followeeID]

	if !exist1 || !exist2 {
		return errors.New("user not found")
	}

	delete(follower.Following, followeeID)
	delete(followee.Followers, followerID)

	return nil
}

func (us *UserService) GetUser(userID int) (*User, error) {
	user, exist := us.Users[userID]
	if !exist {
		return nil, errors.New("user not found")
	}
	return user, nil
}
