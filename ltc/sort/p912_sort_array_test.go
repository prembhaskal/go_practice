package sort

import (
	"fmt"
	"testing"
)

func TestSortArray(t *testing.T) {
	tt := []struct {
		nums []int
	}{
		{
			[]int{5, 2, 3, 1},
		},
		{
			[]int{5, 3, 2, 1},
		},
		{
			[]int{1},
		},
		{[]int{1, 1},},
		{[]int{1, 2},},
		{[]int{1, 2, 1},},
		{[]int{1, 1, 2},},
	}

	for _, tc := range tt {
		fmt.Println()
		fmt.Println()
		fmt.Println("******* TC *********")
		sorted := sortArray(tc.nums)
		fmt.Printf("input array: %v\n", tc.nums)
		fmt.Printf("sorted array: %v\n", sorted)
		fmt.Println("********************")
		fmt.Println()
		fmt.Println()
	}

}
