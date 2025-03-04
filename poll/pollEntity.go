package main

import "time"

type Option struct {
	ID    int
	Text  string
	Votes int
}

type Question struct {
	ID             int
	Text           string
	Options        []Option
	MultipleChoice bool
}
type Poll struct {
	ID          string
	Title       string
	Description string
	Questions   []Question
	CreatedAt   time.Time
	ExpiresAt   time.Time
	Creator     string
}

type Vote struct {
	UserID               string
	QuestionID, OptionID int
}
