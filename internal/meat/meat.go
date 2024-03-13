package meat

import "time"

type MeatType int

// 使用 Enum 定義肉的種類
// 為何使用 iota:
// 1. TypeSafety: 保證不會有其他的值被指派到 MeatType
// 2. Efficiency: Sorting, Searching, Comparing
const (
	Beef MeatType = iota
	Pork
	Chicken
)

type Meat struct {
	Type MeatType
}

const (
	BeefProcessingTime    = 1 * time.Second
	PorkProcessingTime    = 2 * time.Second
	ChickenProcessingTime = 3 * time.Second
)

var ProcessingTimes = map[MeatType]time.Duration{
	Beef:    BeefProcessingTime,
	Pork:    PorkProcessingTime,
	Chicken: ChickenProcessingTime,
}

// 未來可依據 locale 擴充
func (mt MeatType) String() string {
	switch mt {
	case Beef:
		return "Beef"
	case Pork:
		return "Pork"
	case Chicken:
		return "Chicken"
	default:
		return "Unknown"
	}
}
