package service

import (
	"Self/splitwise/models"
	"sync"
)

var (
	onceBalance    sync.Once
	balanceService *BalanceService
)

type BalanceService struct {
	BalanceSheet models.UserToUserBalance
}

func NewBalanceService() *BalanceService {
	onceBalance.Do(func() {
		balanceService = &BalanceService{
			BalanceSheet: models.UserToUserBalance{},
		}
	})
	return balanceService
}

func (bss *BalanceService) AddToSheet(paidBy int, owedBy int, amount float64) {
	if paidBy == owedBy {
		return
	}
	if _, ok := bss.BalanceSheet[paidBy]; !ok {
		bss.BalanceSheet[paidBy] = make(map[int]float64)
	}
	if _, ok := bss.BalanceSheet[owedBy]; !ok {
		bss.BalanceSheet[owedBy] = make(map[int]float64)
	}
	bss.BalanceSheet[paidBy][owedBy] += amount
	bss.BalanceSheet[owedBy][paidBy] -= amount
}

func (bss *BalanceService) GetUserBalances(userID int) map[int]float64 {
	return bss.BalanceSheet[userID]
}
