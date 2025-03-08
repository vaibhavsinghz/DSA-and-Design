package models

type IList interface {
	AddCard(cardID int, name, description string, assigneeID int) error
	RemoveCard(cardID int) error
	UnassignUserFromList(assigneeID int)
	UpdateCardAssigneeFromList(cardID, assigneeID int) error
	UpdateCardStatusFromList(cardID int, status CardStatus) error
}
