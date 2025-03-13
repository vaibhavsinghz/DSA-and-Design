package splitStrategy

type ISplitStrategy interface {
	Split(expense *Expense)
}
