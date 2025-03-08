package models

type ICard interface {
	UpdateStatus(status CardStatus)
	UpdateAssignee(assigneeID int)
	UnassignCard()
	GetAssigneeID() int
}
