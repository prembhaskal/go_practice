package sort

import (
	"fmt"
	"reflect"
	// "runtime"
	"testing"
)

func TestMerge1(t *testing.T) {
	n1 := []int{1, 4, 5, 0, 0, 0}
	n2 := []int{2, 3, 4}

	n3 := []int{1, 2, 3, 4, 4, 5}

	merge1(n1, 3, n2, 3)
	fmt.Printf("%v\n", n1)
	if !reflect.DeepEqual(n1, n3) {
		t.Errorf("exp: %v, act: %v", n3, n1)
		return
	}

	n1 = []int{1, 4, 5, 0, 0, 0}
	n2 = []int{8, 9, 10}
	n3 = []int{1, 4, 5, 8, 9, 10}

	merge1(n1, 3, n2, 3)
	if !reflect.DeepEqual(n3, n1) {
		t.Errorf("exp: %v, act: %v", n3, n1)
		return
	}
	fmt.Printf("%v\n", n1)
}

func TestMerge2(t *testing.T) {
	n1 := []int{1, 4, 5, 0, 0, 0}
	n2 := []int{2, 3, 4}

	merge2(n1, 3, n2, 3)
	fmt.Printf("%v\n", n1)

	n1 = []int{1, 4, 5, 0, 0, 0}
	n2 = []int{8, 9, 10}

	merge2(n1, 3, n2, 3)
	fmt.Printf("%v\n", n1)
}
