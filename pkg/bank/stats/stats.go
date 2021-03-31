package stats

import "github.com/AKozikhodzha/bank/pkg/bank/types"

//Avg
func Avg(payments []types.Payment) types.Money {
	sum := types.Money(0)
	for _, card := range payments {
		sum += card.Amount
	}
	middleValue := sum / types.Money(len(payments))
	return middleValue
}
