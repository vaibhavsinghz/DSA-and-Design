package splitStrategy

type EqualStrategy struct {
}

func (e EqualStrategy) Split(expense *Expense) {
	noOfUserInvolved := len(expense.SplitBetween)

	equalSplitVal := expense.Amount / float64(noOfUserInvolved)

	for i := 0; i < noOfUserInvolved; i++ {
		expense.SplitBetween[i].Amount = equalSplitVal
	}
}

func NewEqualStrategy() ISplitStrategy {
	return &EqualStrategy{}
}
