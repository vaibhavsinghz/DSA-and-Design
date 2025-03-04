package main

import "sort"

type FeedService struct {
	UserQuery  UserQuery
	PostReader PostReader
}

func NewFeedService(userQuery UserQuery, postReader PostReader) *FeedService {
	return &FeedService{UserQuery: userQuery, PostReader: postReader}
}

func (fs *FeedService) GetFeed(userID, pageSize, lastPostID int) Feed {
	user, err := fs.UserQuery.GetUser(userID)
	if err != nil {
		return Feed{}
	}

	feedPostIDs := []int{}
	userPostIDs, err := fs.PostReader.GetUserPostIDs(userID)
	if err != nil {
		return Feed{}
	}
	feedPostIDs = append(feedPostIDs, userPostIDs...)

	for followerID := range user.Following {
		followerPostIDs, err := fs.PostReader.GetUserPostIDs(followerID)
		if err != nil {
			return Feed{}
		}
		feedPostIDs = append(feedPostIDs, followerPostIDs...)
	}

	feedPosts := []Post{}
	for _, postID := range feedPostIDs {
		post, err := fs.PostReader.GetPost(postID)
		if err != nil {
			return Feed{}
		}
		feedPosts = append(feedPosts, post)
	}

	//sort Post based on time
	sort.Slice(feedPosts, func(i, j int) bool {
		return feedPosts[i].Timestamp.After(feedPosts[j].Timestamp)
	})

	//pagination
	start := 0
	if lastPostID > 0 {
		for i, post := range feedPosts {
			if post.ID == lastPostID {
				start = i + 1
				break
			}
		}
	}

	end := min(start+pageSize, len(feedPosts))

	nextPage := 0
	if end < len(feedPosts) {
		nextPage = feedPosts[end-1].ID
	}

	return Feed{
		Posts:    feedPosts[start:end],
		NextPage: nextPage,
	}
}
