package main

import "time"

type User struct {
	ID                   int
	Name, Username       string
	Followers, Following map[int]struct{}
}

type Post struct {
	ID, UserID int
	Content    string
	Timestamp  time.Time
}

type Feed struct {
	Posts    []Post
	NextPage int
}
