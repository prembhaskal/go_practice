package generics_test

import (
	"fmt"
	"testing"

	generics "github.com/prembhaskal/go_practice/pkg/learning/go/generics"
)

func TestLinkedList(t *testing.T) {
	vals := []int{3, 4, 9, 0, 1}
	head := generics.NewLinkedList(vals)
	fmt.Printf("head: %v\n", head)

	vals2 := []int{3, 0, 8, 5, 1, 9}
	head2 := generics.NewLinkedList(vals2)
	fmt.Printf("head2: %v\n", head2)
}
