package main

import (
	"errors"
	"sync"
	"time"
)

type PostService struct {
	Posts         map[int]Post
	UserPosts     map[int][]int //UserID -> List of PostIDs (for quick access)
	postIDCounter int
	mu            sync.Mutex
}

func NewPostService() *PostService {
	return &PostService{
		Posts:         map[int]Post{},
		UserPosts:     map[int][]int{},
		postIDCounter: 0,
	}
}

func (ps *PostService) CreatePost(userID int, content string) Post {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	ps.postIDCounter++
	newPost := Post{
		ID:        ps.postIDCounter,
		UserID:    userID,
		Content:   content,
		Timestamp: time.Now(),
	}

	ps.Posts[newPost.ID] = newPost
	ps.UserPosts[userID] = append(ps.UserPosts[userID], newPost.ID)
	return newPost
}

func (ps *PostService) DeletePost(userID, postID int) error {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	if _, exist := ps.Posts[postID]; !exist {
		return errors.New("post not found")
	}

	delete(ps.Posts, postID)
	for i, userPost := range ps.UserPosts[userID] {
		if userPost == postID {
			ps.UserPosts[userID] = append(ps.UserPosts[userID][:i], ps.UserPosts[userID][i+1:]...)
			break
		}
	}
	return nil
}

func (ps *PostService) GetPost(postID int) (Post, error) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	post, exist := ps.Posts[postID]
	if !exist {
		return Post{}, errors.New("post does not exist")
	}

	return post, nil
}

func (ps *PostService) GetUserPostIDs(userID int) ([]int, error) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	posts, exist := ps.UserPosts[userID]
	if !exist {
		return []int{}, errors.New("user posts does not exist")
	}

	return posts, nil
}
