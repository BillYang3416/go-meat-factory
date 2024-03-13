package supplier

import "github.com/billyang3416/go-meat-factory/internal/meat"

func SupplyMeat(totalMeat, numOfBeef, numOfPork, numOfChicken int) []meat.Meat {

	meats := make([]meat.Meat, totalMeat)

	for i := 0; i < totalMeat; i++ {
		switch {
		case i < numOfBeef:
			meats = append(meats, meat.Meat{ID: i, Type: meat.Beef})
		case i < numOfBeef+numOfPork:
			meats = append(meats, meat.Meat{ID: i, Type: meat.Pork})
		default:
			meats = append(meats, meat.Meat{ID: i, Type: meat.Chicken})
		}
	}

	return meats
}
