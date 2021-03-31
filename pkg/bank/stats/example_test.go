package stats

import (
	"fmt"

	"github.com/AKozikhodzha/bank/pkg/bank/types"
)

func ExampleAvg() {
	payment := []types.Payment{
		{
			ID:     5555,
			Amount: 1_000_00,
		},
		{
			ID:     5555,
			Amount: 2_000_00,
		},
		{
			ID:     5555,
			Amount: 3_000_00,
		},
		{
			ID:     5555,
			Amount: 4_000_00,
		},
	}
	fmt.Println(Avg(payment))
	//Output:250000

}
