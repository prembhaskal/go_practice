package array

import (
	"fmt"
	"testing"
)

func TestThirdLarge(t *testing.T) {
	tbl := [][]int{
		{2, 4, 3, 1},
		{2, 2, 1, 1},
		{2, 2, 3, 1},
		{2, 2, 3, 2},
	}

	for _, ar := range tbl {
		fmt.Printf("array is %v\n", ar)
		res := thirdMax(ar)
		fmt.Printf("result is %d\n", res)
		fmt.Println()
	}
}
