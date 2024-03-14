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
	ID   int // 用來辨識沒有員工取到相同的肉
	Type MeatType
}

const (
	BeefProcessingTime    = 1 * time.Second
	PorkProcessingTime    = 2 * time.Second
	ChickenProcessingTime = 3 * time.Second
)

var processingTimes = map[MeatType]time.Duration{
	Beef:    BeefProcessingTime,
	Pork:    PorkProcessingTime,
	Chicken: ChickenProcessingTime,
}

// 未來可依據 locale 擴充
func (mt MeatType) String() string {
	switch mt {
	case Beef:
		return "牛肉"
	case Pork:
		return "豬肉"
	case Chicken:
		return "雞肉"
	default:
		return "不明的肉類"
	}
}

func ProcessingTime(mt MeatType) time.Duration {
	return processingTimes[mt]
}
