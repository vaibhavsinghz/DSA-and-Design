package models

import (
	"Self/splitwise/models/splitStrategy"
	"fmt"
)

type Group struct {
	ID      int
	Name    string
	Users   map[int]struct{} //list of userID
	Expense []*splitStrategy.Expense
}

func NewGroup(id int, name string, users []int) *Group {
	group := &Group{
		ID:      id,
		Name:    name,
		Expense: []*splitStrategy.Expense{},
		Users:   map[int]struct{}{},
	}

	for _, userID := range users {
		group.Users[userID] = struct{}{}
	}

	return group
}

func (group *Group) AddUser(userID int) {
	group.Users[userID] = struct{}{}
}

func (group *Group) RemoveUser(userID int) {
	delete(group.Users, userID)
}

func (group *Group) AddExpense(expense *splitStrategy.Expense) {
	group.Expense = append(group.Expense, expense)
}

func (group *Group) VerifyExpense(expense *splitStrategy.Expense) error {
	paidByUserID := expense.PaidBy
	if _, exists := group.Users[paidByUserID]; !exists {
		return fmt.Errorf("user paid does not belong to group %d", group.ID)
	}

	var currentTotalSplitAmount float64

	// we don't need to verify split list for splitType as equal
	if expense.SplitType == splitStrategy.EQUAL {
		var splitList []*splitStrategy.Split
		for userID, _ := range group.Users {
			splitList = append(splitList, splitStrategy.NewSplit(userID, 0.0))
		}
		expense.SplitBetween = splitList
		return nil
	}

	//for other case user will specify the amount and user that we need to verify from existing group
	for _, split := range expense.SplitBetween {
		if split == nil {
			return fmt.Errorf("invalid expense split list")
		}
		if _, exists := group.Users[split.UserID]; !exists {
			return fmt.Errorf("user %d does not belong to group %d", split.UserID, group.ID)
		}
		currentTotalSplitAmount += split.Amount
	}
	//validate logic need to be in strategy
	if expense.SplitType == splitStrategy.EQUAL && currentTotalSplitAmount != expense.Amount {
		return fmt.Errorf("expense split amount does not match amount of expense")
	}
	if expense.SplitType == splitStrategy.PERCENT && currentTotalSplitAmount != 100 {
		return fmt.Errorf("expense split amount does not match amount of expense")
	}

	return nil
}
