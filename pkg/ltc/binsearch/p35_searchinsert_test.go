package binsearch

import (
	"testing"
)

func TestSearchInsert1(t *testing.T) {
	ar := []int{1, 3, 5, 6, 7, 9, 10, 11}
	trgt := 8

	exp := 5

	testInsertMethod(t, ar, trgt, exp)

	testInsertMethod(t, []int{1, 3, 5, 6}, 5, 2)
	testInsertMethod(t, []int{1, 3, 5, 6}, 7, 4)
	testInsertMethod(t, []int{1, 3, 5, 6}, 0, 0)

	// act := searchInsert(ar, trgt)

	// if exp != act {
	// 	t.Errorf("exp: %d, act: %d", exp, act)
	// }
}

func testInsertMethod(t *testing.T, ar []int, trgt, exp int) {
	// ar := []int{1, 3, 5, 6, 7, 9, 10, 11}
	// trgt := 7

	// exp := 4
	act := searchInsert(ar, trgt)

	if exp != act {
		t.Errorf("exp: %d, act: %d", exp, act)
	}
}
