package binsearch

import (
	"testing"
)

func TestSearch(t *testing.T) {
	arr := []int{1, 2, 5, 6, 10, 13, 15}

	var find int

	find = 0
	idx := search(arr, find)

	assertInt(t, idx, -1)

	for idx, val := range arr {
		act := search(arr, val)
		assertInt(t, act, idx)
	}

	arr = []int{1}

	assertInt(t, search(arr, 0), -1)
	assertInt(t, search(arr, 1), 0)

	arr = []int{1, 10}

	assertInt(t, search(arr, 1), 0)
	assertInt(t, search(arr, 10), 1)
}

func assertInt(t *testing.T, act, exp int) {
	if act != exp {
		t.Errorf("exp: %d, act: %d", exp, act)
	}
}
