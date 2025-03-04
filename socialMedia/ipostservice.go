package main

type PostWriter interface {
	CreatePost(userID int, content string) Post
	DeletePost(userID, postID int) error
}

type PostReader interface {
	GetPost(postID int) (Post, error)
	GetUserPostIDs(userID int) ([]int, error)
}
