package avg

import (
	"Avg_map/types"
)

func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	categories := map[types.Category]types.Money{}
	counts := map[types.Category]types.Money{}
	//var result types.Money
	for _, payment := range payments {
		categories[payment.Category] += payment.Amount
		counts[payment.Category] += 1
	}

	for key := range categories {
		categories[key] /= counts[key]
	}

	return categories
}
