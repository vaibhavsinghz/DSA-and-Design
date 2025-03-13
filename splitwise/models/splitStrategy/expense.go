package splitStrategy

type Expense struct {
	Description   string
	Amount        float64
	PaidBy        int // userID
	SplitBetween  []*Split
	SplitType     SplitType
	SplitStrategy ISplitStrategy
}

func (expense *Expense) Split() {
	expense.SplitStrategy.Split(expense)
}

func (expense *Expense) GetSplitList() []*Split {
	return expense.SplitBetween
}
