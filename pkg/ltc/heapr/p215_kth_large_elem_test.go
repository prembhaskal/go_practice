package heapr

import (
	"fmt"
	"testing"
)

func TestNewMinHeap(t *testing.T) {
	ar := make([]int, 0)

	for i := 20; i >= 0; i-- {
		ar = append(ar, i)
	}

	hp := newintminheap(ar)
	fmt.Printf("heap content: %s\n", hp)

	for i := 0; i < len(ar); i++ {
		min := hp.extractMin()
		fmt.Printf("min: %d\n", min)
	}
}

func TestNewMaxHeap(t *testing.T) {
	ar := make([]int, 0)

	for i := 0; i < 21; i++ {
		ar = append(ar, i)
	}

	hp := newintmaxheap(ar)
	fmt.Printf("max heap content: %s\n", hp)

	for i := 0; i < len(ar); i++ {
		min := hp.extractMax()
		fmt.Printf("max: %d\n", min)
	}
}
