package binsearchrot

import (
	"testing"

	"github.com/prembhaskal/go_practice/util"
)

func TestRotateSearch(t *testing.T) {
	arr := []int{9, 11, 12, 15, -2, -1, 0, 5}
	var find int
	find = -5

	util.AssertInt(t, -1, search(arr, find))
	for k, v := range arr {
		util.AssertInt(t, k, search(arr, v))
	}

	arr = []int{-1, 0, 1, 2, 4, 5, 10}
	for k, v := range arr {
		util.AssertInt(t, k, search(arr, v))
	}
}
