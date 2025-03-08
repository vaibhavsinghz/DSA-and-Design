package models

type Card struct {
	ID          int
	Name        string
	Description string
	AssigneeID  int
	Status      CardStatus
}

func NewCard(cardID int, name, description string, assigneeID int) ICard {
	return &Card{
		ID:          cardID,
		Name:        name,
		Description: description,
		AssigneeID:  assigneeID,
		Status:      ToDo,
	}
}

func (card *Card) UpdateStatus(status CardStatus) {
	card.Status = status
}

func (card *Card) UpdateAssignee(assigneeID int) {
	card.AssigneeID = assigneeID
}

func (card *Card) UnassignCard() {
	card.AssigneeID = 0
}

func (card *Card) GetAssigneeID() int {
	return card.AssigneeID
}
