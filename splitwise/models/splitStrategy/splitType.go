package splitStrategy

type SplitType int

const (
	EQUAL SplitType = iota
	FIXED
	PERCENT
)

//func (et SplitType) String() string {
//	return [...]string{"EQUAL", "FIXED", "PERCENT"}[et]
//}
