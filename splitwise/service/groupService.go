package service

import (
	"Self/splitwise/models"
	"Self/splitwise/models/splitStrategy"
	"fmt"
	"sync"
)

type GroupService struct {
	GroupMap       map[int]*models.Group
	nextGroupId    int
	UserService    *UserService
	BalanceService *BalanceService
}

var (
	groupService *GroupService
	onceGroup    sync.Once
)

func NewGroupService() *GroupService {
	onceGroup.Do(func() {
		groupService = &GroupService{
			GroupMap:       make(map[int]*models.Group),
			nextGroupId:    1,
			UserService:    NewUserService(),
			BalanceService: NewBalanceService(),
		}
	})
	return groupService
}

func (groupService *GroupService) GetGroupById(groupId int) *models.Group {
	return groupService.GroupMap[groupId]
}

func (groupService *GroupService) CreateGroup(name string, userIDs []int) int {
	groupID := groupService.nextGroupId

	validUserIDs := make([]int, 0, len(userIDs))

	for _, userID := range userIDs {
		user := groupService.UserService.GetUserById(userID)
		if user == nil {
			fmt.Printf("%d userID does not exist\n", userID)
			continue
		}
		validUserIDs = append(validUserIDs, userID)
	}

	group := models.NewGroup(groupID, name, validUserIDs)
	groupService.GroupMap[groupID] = group

	groupService.nextGroupId++
	return groupID
}

func (groupService *GroupService) AddUserToGroup(groupID, userID int) error {
	group := groupService.GroupMap[groupID]
	if group == nil {
		return fmt.Errorf("group %d does not exist", groupID)
	}
	user := groupService.UserService.GetUserById(userID)
	if user == nil {
		return fmt.Errorf("%d userID does not exist\n", userID)
	}
	group.AddUser(user.ID)
	return nil
}

func (groupService *GroupService) RemoveUserFromGroup(groupID, userID int) error {
	group := groupService.GroupMap[groupID]
	if group == nil {
		return fmt.Errorf("group %d does not exist", groupID)
	}
	user := groupService.UserService.GetUserById(userID)
	if user == nil {
		return fmt.Errorf("%d userID does not exist\n", userID)
	}
	group.RemoveUser(user.ID)
	return nil
}

func (groupService *GroupService) AddExpense(groupID int, description string, amount float64, paidBy int, splitType splitStrategy.SplitType, splitBetween []*splitStrategy.Split) error {
	group, exist := groupService.GroupMap[groupID]
	if !exist {
		return fmt.Errorf("group %d does not exist", groupID)
	}

	expense := NewExpense(paidBy, amount, splitBetween, splitType, description)

	if err := group.VerifyExpense(expense); err != nil {
		return err
	}

	expense.Split()
	group.AddExpense(expense)

	for _, split := range expense.GetSplitList() {
		if split == nil {
			fmt.Println("invalid expense split")
			continue
		}
		groupService.BalanceService.AddToSheet(paidBy, split.UserID, split.Amount)
	}

	return nil
}

func (groupService *GroupService) GetUserBalanceRecord(userID int) map[int]float64 {
	balance := groupService.BalanceService.GetUserBalances(userID)
	for relatedUserID, amt := range balance {
		if amt > 0 {
			fmt.Printf("user %d owes %.2f amt to %d\n", userID, amt, relatedUserID)
		} else {
			fmt.Printf("user %d has to pay %.2f amt to %d\n", userID, -amt, relatedUserID)
		}
	}
	return balance
}
