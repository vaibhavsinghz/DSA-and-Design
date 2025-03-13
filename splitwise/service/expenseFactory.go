package service

import (
	"Self/splitwise/models/splitStrategy"
)

func NewExpense(paidBy int, amount float64, splitBetween []*splitStrategy.Split, splitType splitStrategy.SplitType, description string) *splitStrategy.Expense {
	switch splitType {
	case splitStrategy.EQUAL:
		return &splitStrategy.Expense{Description: description, Amount: amount, PaidBy: paidBy, SplitBetween: splitBetween, SplitType: splitType, SplitStrategy: splitStrategy.NewEqualStrategy()}
	case splitStrategy.PERCENT:
		return &splitStrategy.Expense{Description: description, Amount: amount, PaidBy: paidBy, SplitBetween: splitBetween, SplitType: splitType, SplitStrategy: splitStrategy.NewPercentageStrategy()}
	case splitStrategy.FIXED:
		return &splitStrategy.Expense{Description: description, Amount: amount, PaidBy: paidBy, SplitBetween: splitBetween, SplitType: splitType, SplitStrategy: splitStrategy.NewFixedStrategy()}
	default:
		return nil
	}
}

//3,05,24,000
