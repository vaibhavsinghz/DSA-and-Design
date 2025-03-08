package models

import "fmt"

type privacyState string

const (
	//privacy status
	Public  privacyState = "public"
	Private privacyState = "private"

	//URL
	urlFormat = "https://www.jira.io/%d/%s" // /boardID/name
)

type Board struct {
	ID      int
	Name    string
	Privacy privacyState
	URL     string
	Members map[int]struct{} // set of user id
	Lists   map[int]IList    //map of listID to list
}

func NewBoard(id int, name string, privacyState privacyState, users []int) IBoard {
	board := &Board{
		ID:      id,
		Name:    name,
		Privacy: privacyState,
		URL:     fmt.Sprintf(urlFormat, id, name),
		Members: make(map[int]struct{}),
		Lists:   make(map[int]IList),
	}
	board.AddMembers(users)
	return board
}

func (board *Board) GetBoardID() int {
	return board.ID
}

func (board *Board) GetName() string {
	return board.Name
}

func (board *Board) AddMembers(memberIDs []int) {
	for _, memberID := range memberIDs {
		board.Members[memberID] = struct{}{}
	}
}

func (board *Board) RemoveMembers(memberIDs []int) {
	for _, memberID := range memberIDs {
		delete(board.Members, memberID)
		for _, list := range board.Lists {
			list.UnassignUserFromList(memberID)
		}
	}
}

func (board *Board) AddList(listID int, name string) error {
	if _, exists := board.Lists[listID]; exists {
		return fmt.Errorf("list %d already exists", listID)
	}

	list := NewList(listID, name)
	board.Lists[listID] = list

	return nil
}

func (board *Board) RemoveList(listID int) error {
	if _, exists := board.Lists[listID]; !exists {
		return fmt.Errorf("list %d does not exists", listID)
	}

	delete(board.Lists, listID)
	return nil
}

func (board *Board) AddCardInBoardList(cardID, listID int, cardName, description string, assigneeID int) error {
	cardID = -1
	list, exist := board.Lists[listID]
	if !exist {
		return fmt.Errorf("list %d does not exist", listID)
	}
	if err := list.AddCard(cardID, cardName, description, assigneeID); err != nil {
		return err
	}
	return nil
}

func (board *Board) RemoveCardFromBoardList(listID, cardID int) error {
	list, exist := board.Lists[listID]
	if !exist {
		return fmt.Errorf("list %d does not exist", listID)
	}
	if err := list.RemoveCard(cardID); err != nil {
		return err
	}
	return nil
}

func (board *Board) UpdateCardAssigneeFromBoardList(listID, cardID, assigneeID int) error {
	list, exist := board.Lists[listID]
	if !exist {
		return fmt.Errorf("list %d does not exist", listID)
	}

	return list.UpdateCardAssigneeFromList(cardID, assigneeID)
}

func (board *Board) UpdateCardStatusFromBoardList(listID, cardID int, status CardStatus) error {
	list, exist := board.Lists[listID]
	if !exist {
		return fmt.Errorf("list %d does not exist", listID)
	}

	return list.UpdateCardStatusFromList(cardID, status)
}
