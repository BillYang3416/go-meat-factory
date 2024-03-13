package supplier

import "github.com/billyang3416/go-meat-factory/internal/meat"

func SupplyMeat(totalMeat, numOfBeef, numOfPork, numOfChicken int) []meat.Meat {

	meats := make([]meat.Meat, totalMeat)

	for i := 0; i < totalMeat; i++ {
		ID := i + 1
		switch {
		case i < numOfBeef:
			meats[i] = meat.Meat{ID: ID, Type: meat.Beef}
		case i < numOfBeef+numOfPork:
			meats[i] = meat.Meat{ID: ID, Type: meat.Pork}
		default:
			meats[i] = meat.Meat{ID: ID, Type: meat.Chicken}
		}
	}

	return meats
}
