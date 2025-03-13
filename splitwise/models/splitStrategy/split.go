package splitStrategy

type Split struct {
	UserID int
	Amount float64
}

func NewSplit(userID int, amount float64) *Split {
	return &Split{
		UserID: userID,
		Amount: amount,
	}
}

func NewSplitList(splitMap map[int]float64) []*Split {
	var splitList []*Split
	for userID, amount := range splitMap {
		splitList = append(splitList, NewSplit(userID, amount))
	}
	return splitList
}
