package main

import (
	"errors"
	"sync"
	"time"
)

type PollService struct {
	polls map[string]*Poll
	votes map[string][]Vote
	mu    sync.RWMutex
}

func NewPollService() *PollService {
	return &PollService{
		polls: make(map[string]*Poll),
		votes: make(map[string][]Vote),
	}
}

func (ps *PollService) CreatePoll(pollId, title, description, creator string, questions []Question, expiryDuration time.Duration) (*Poll, error) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	if _, exist := ps.polls[pollId]; exist {
		return nil, errors.New("poll already exists")
	}

	for i := range questions {
		questions[i].ID = i + 1

		for j := range questions[i].Options {
			questions[i].Options[j].ID = j + 1
		}
	}

	newPoll := &Poll{
		ID:          pollId,
		Title:       title,
		Description: description,
		CreatedAt:   time.Now(),
		ExpiresAt:   time.Now().Add(expiryDuration),
		Creator:     creator,
		Questions:   questions,
	}

	ps.polls[pollId] = newPoll
	ps.votes[pollId] = []Vote{}

	return newPoll, nil
}

func (ps *PollService) VoteOnPoll(pollId, userID string, questionID, optionID int) error {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	poll, exist := ps.polls[pollId]
	if !exist {
		return errors.New("poll does not exist")
	}

	if time.Now().After(poll.ExpiresAt) {
		return errors.New("poll expired")
	}

	var question *Question
	for i := range poll.Questions {
		if poll.Questions[i].ID == questionID {
			question = &poll.Questions[i]
			break
		}
	}

	if question == nil {
		return errors.New("question does not exist")
	}

	var option *Option
	for i := range question.Options {
		if question.Options[i].ID == optionID {
			option = &question.Options[i]
			break
		}
	}

	if option == nil {
		return errors.New("option does not exist")
	}

	for _, vote := range ps.votes[pollId] {
		if vote.UserID == userID && vote.QuestionID == questionID {
			if !question.MultipleChoice || vote.OptionID == optionID {
				return errors.New("vote already present")
			}
		}
	}

	newVote := Vote{
		UserID:     userID,
		QuestionID: questionID,
		OptionID:   optionID,
	}

	ps.votes[pollId] = append(ps.votes[pollId], newVote)
	option.Votes++
	return nil
}

func (ps *PollService) UpdatePollQuestion(pollId string, questionID int, questionText string, newOptions []Option, multipleChoice bool) error {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	poll, exist := ps.polls[pollId]
	if !exist {
		return errors.New("poll does not exist")
	}

	var question *Question
	for i := range poll.Questions {
		if poll.Questions[i].ID == questionID {
			question = &poll.Questions[i]
			break
		}
	}

	if question == nil {
		return errors.New("question does not exist")
	}

	question.Text = questionText
	question.MultipleChoice = multipleChoice

	if newOptions != nil {
		for i := range newOptions {
			newOptions[i].ID = i + 1
		}
		question.Options = newOptions
	}

	return nil
}

func (ps *PollService) GetPoll(pollId string) (*Poll, error) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	poll, exist := ps.polls[pollId]
	if !exist {
		return nil, errors.New("poll does not exist")
	}

	return poll, nil
}

func (ps *PollService) DeletePoll(pollId string) error {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	if _, exist := ps.polls[pollId]; !exist {
		return errors.New("poll does not exist")
	}

	delete(ps.polls, pollId)
	delete(ps.votes, pollId)

	return nil
}
