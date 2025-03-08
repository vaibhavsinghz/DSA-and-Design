package models

import "fmt"

type List struct {
	ID    int
	Name  string
	Cards map[int]ICard //card id to card map
}

func NewList(listID int, name string) IList {
	return &List{
		ID:    listID,
		Name:  name,
		Cards: make(map[int]ICard),
	}
}

func (list *List) AddCard(cardID int, name, description string, assigneeID int) error {
	if _, exist := list.Cards[cardID]; exist {
		return fmt.Errorf("card %d already exists", cardID)
	}

	card := NewCard(cardID, name, description, assigneeID)
	list.Cards[cardID] = card

	return nil
}

func (list *List) RemoveCard(cardID int) error {
	if _, exist := list.Cards[cardID]; !exist {
		return fmt.Errorf("card %d does not exist", cardID)
	}
	delete(list.Cards, cardID)
	return nil
}

func (list *List) UnassignUserFromList(assigneeID int) {
	for _, card := range list.Cards {
		if card.GetAssigneeID() == assigneeID {
			card.UnassignCard()
		}
	}
}

func (list *List) UpdateCardAssigneeFromList(cardID, assigneeID int) error {
	card, exist := list.Cards[cardID]
	if !exist {
		return fmt.Errorf("card %d does not exist", cardID)
	}

	card.UpdateAssignee(assigneeID)
	return nil
}

func (list *List) UpdateCardStatusFromList(cardID int, status CardStatus) error {
	card, exist := list.Cards[cardID]
	if !exist {
		return fmt.Errorf("card %d does not exist", cardID)
	}

	card.UpdateStatus(status)
	return nil
}
