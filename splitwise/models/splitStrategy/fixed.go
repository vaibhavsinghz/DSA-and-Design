package splitStrategy

type FixedStrategy struct {
}

func (f FixedStrategy) Split(expense *Expense) {
	return
}

func NewFixedStrategy() ISplitStrategy {
	return &FixedStrategy{}
}
