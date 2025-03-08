package models

type IBoard interface {
	GetBoardID() int
	GetName() string
	AddMembers(memberIDs []int)
	RemoveMembers(memberIDs []int)
	AddList(listID int, name string) error
	RemoveList(listID int) error
	AddCardInBoardList(cardID, listID int, cardName, description string, assigneeID int) error
	RemoveCardFromBoardList(listID, cardID int) error
	UpdateCardAssigneeFromBoardList(listID, cardID, assigneeID int) error
	UpdateCardStatusFromBoardList(listID, cardID int, status CardStatus) error
}
