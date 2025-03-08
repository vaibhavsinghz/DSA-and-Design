package models

type CardStatus int

const (
	ToDo = iota
	InProgress
	DevDone
	CodeReview
	QAPending
	QAStarted
	QADone
	Live
	Blocked
)
