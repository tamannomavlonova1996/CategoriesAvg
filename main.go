package main

import (
	"Avg_map/avg"
	"Avg_map/types"
	"fmt"
)

func main() {
	payments := []types.Payment{
		{
			ID:       2,
			Amount:   20,
			Category: "wallet",
		},
		{
			ID:       1,
			Amount:   30,
			Category: "Bonus",
		},
		{
			ID:       3,
			Amount:   40,
			Category: "Bonus",
		},
	}
	fmt.Println(avg.CategoriesAvg(payments))
}
