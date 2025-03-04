package main

type UserRelationManager interface {
	Follow(followerID, followeeID int) error
	UnFollow(followerID, followeeID int) error
}

type UserQuery interface {
	GetUser(userID int) (*User, error)
}
