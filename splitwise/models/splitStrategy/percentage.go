package splitStrategy

type PercentageStrategy struct {
}

func (p PercentageStrategy) Split(expense *Expense) {
	noOfUserInvolved := len(expense.SplitBetween)

	for i := 0; i < noOfUserInvolved; i++ {
		percentage := expense.SplitBetween[i].Amount
		percentageValue := (percentage * expense.Amount) / 100
		expense.SplitBetween[i].Amount = percentageValue
	}
}

func NewPercentageStrategy() ISplitStrategy {
	return &PercentageStrategy{}
}
