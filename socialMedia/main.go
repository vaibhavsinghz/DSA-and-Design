package main

import "fmt"

func main() {
	userService := NewUserService()
	postService := NewPostService()
	feedService := NewFeedService(userService, postService)

	// Create Users
	userService.Users[1] = &User{ID: 1, Name: "Alice", Username: "alice", Followers: make(map[int]struct{}), Following: make(map[int]struct{})}
	userService.Users[2] = &User{ID: 2, Name: "Bob", Username: "bob", Followers: make(map[int]struct{}), Following: make(map[int]struct{})}

	// Follow User
	userService.Follow(1, 2)

	// Create Posts
	postService.CreatePost(1, "Hello from Alice!")
	postService.CreatePost(2, "Hello from Bob!")

	// Get Feed
	feed := feedService.GetFeed(1, 10, 0)
	fmt.Println(feed)
}
