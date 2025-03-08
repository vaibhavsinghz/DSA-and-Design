package models

import (
	"errors"
	"fmt"
)

type JIRA struct {
	Users          map[int]*User  //user id to use map
	Boards         map[int]IBoard //board id to board map
	boardIDCounter int
	userIDCounter  int
	listIDCounter  int
	cardIDCounter  int
}

func InitJIRA() IJira {
	return &JIRA{
		Boards: make(map[int]IBoard),
	}
}

func (jira *JIRA) AddBoard(name string, isPrivate bool, userIDs []int) (boardID int, err error) {
	privacy := Public
	if isPrivate {
		privacy = Private
	}

	if !jira.verifyExistingUser(userIDs) {
		err = errors.New("user does not exist")
	}

	jira.boardIDCounter++
	board := NewBoard(jira.boardIDCounter, name, privacy, userIDs)
	boardID = board.GetBoardID()

	jira.Boards[boardID] = board
	return
}

func (jira *JIRA) RemoveBoard(boardID int) (err error) {
	if _, exist := jira.Boards[boardID]; !exist {
		err = errors.New("board does not exist")
		return
	}
	delete(jira.Boards, boardID)
	return
}

func (jira *JIRA) verifyExistingUser(userIDs []int) bool {
	for _, userID := range userIDs {
		if _, exist := jira.Users[userID]; !exist {
			return false
		}
	}
	return true
}

func (jira *JIRA) AddUser(name, email string, boardIDs []int) (userID int) {
	jira.userIDCounter++
	user := NewUser(jira.userIDCounter, name, email)
	userID = user.ID

	for _, boardID := range boardIDs {
		board, exist := jira.Boards[boardID]
		if !exist {
			fmt.Printf("%d boardID does not exist\n", boardID)
		}
		board.AddMembers([]int{user.ID})
		fmt.Printf("Added user %s to board %s\n", user.Name, board.GetName())
	}

	jira.Users[user.ID] = user
	return
}

func (jira *JIRA) RemoveUser(userID int) {
	if _, exist := jira.Users[userID]; !exist {
		return
	}

	delete(jira.Users, userID)

	for _, board := range jira.Boards {
		board.RemoveMembers([]int{userID})
	}
}

func (jira *JIRA) AddListToBoard(boardID int, listName string) (listID int, err error) {
	listID = -1
	board, exist := jira.Boards[boardID]
	if !exist {
		err = fmt.Errorf("board %d does not exist", boardID)
		return
	}

	jira.listIDCounter++
	err = board.AddList(jira.listIDCounter, listName)
	if err == nil {
		listID = jira.listIDCounter
	}
	return
}

func (jira *JIRA) RemoveListFromBoard(boardID, listID int) (err error) {
	board, exist := jira.Boards[boardID]
	if !exist {
		err = fmt.Errorf("board %d does not exist", boardID)
		return
	}

	return board.RemoveList(listID)
}

func (jira *JIRA) AddCardInBoardList(boardID, listID int, cardName, description string, assigneeID int) (cardID int, err error) {
	cardID = -1
	board, exist := jira.Boards[boardID]
	if !exist {
		err = fmt.Errorf("board %d does not exist", boardID)
		return
	}

	if _, exist := jira.Users[assigneeID]; !exist {
		assigneeID = 0
		fmt.Printf("user %d does not exist, assignment skipped", assigneeID)
	}

	jira.cardIDCounter++
	err = board.AddCardInBoardList(jira.cardIDCounter, listID, cardName, description, assigneeID)
	if err == nil {
		cardID = jira.cardIDCounter
	}
	return
}

func (jira *JIRA) RemoveCardInBoardList(boardID, listID, cardID int) (err error) {
	board, exist := jira.Boards[boardID]
	if !exist {
		err = fmt.Errorf("board %d does not exist", boardID)
		return
	}

	err = board.RemoveCardFromBoardList(listID, cardID)

	return
}

func (jira *JIRA) UpdateCardAssignee(boardID, listID, cardID int, assigneeID int) (err error) {
	_, exist := jira.Users[assigneeID]
	if !exist {
		err = fmt.Errorf("user %d does not exist", assigneeID)
	}

	board, exist := jira.Boards[boardID]
	if !exist {
		err = fmt.Errorf("board %d does not exist", boardID)
		return
	}

	err = board.UpdateCardAssigneeFromBoardList(listID, cardID, assigneeID)
	return
}

func (jira *JIRA) UpdateCardStatus(boardID, listID, cardID int, status CardStatus) (err error) {
	board, exist := jira.Boards[boardID]
	if !exist {
		err = fmt.Errorf("board %d does not exist", boardID)
		return
	}

	err = board.UpdateCardStatusFromBoardList(listID, cardID, status)
	return
}
