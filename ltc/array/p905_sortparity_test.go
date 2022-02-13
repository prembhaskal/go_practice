package array

import (
	"fmt"
	"testing"
)

func TestSortParity(t *testing.T) {
	tbl := [][]int{
		{1, 2, 3, 4},
		{2, 1, 3, 4},
		{0},
		{2, 4, 6},
		{1, 3, 5},
	}

	for _, ar := range tbl {
		fmt.Printf("before: %v\n", ar)
		res := sortArrayByParity(ar)
		fmt.Printf("after: %v\n", res)
	}
}
