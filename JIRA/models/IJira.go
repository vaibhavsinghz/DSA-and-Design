package models

type IJira interface {
	AddBoard(name string, isPrivate bool, userIDs []int) (boardID int, err error)
	RemoveBoard(boardID int) (err error)
	AddUser(name, email string, boardIDs []int) (userID int)
	RemoveUser(userID int)
	AddListToBoard(boardID int, listName string) (listID int, err error)
	RemoveListFromBoard(boardID, listID int) (err error)
	AddCardInBoardList(boardID, listID int, cardName, description string, assigneeID int) (cardID int, err error)
	RemoveCardInBoardList(boardID, listID, cardID int) (err error)
	UpdateCardAssignee(boardID, listID, cardID int, assigneeID int) (err error)
	UpdateCardStatus(boardID, listID, cardID int, status CardStatus) (err error)
}
