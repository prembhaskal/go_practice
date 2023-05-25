package generics

import (
	"fmt"
	"strings"
)

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func NewLinkedList[T any](vals []T) *List[T] {
	var prev *List[T]
	var head *List[T]
	for _, v := range vals {
		curr := &List[T]{
			val: v,
		}
		if prev != nil {
			prev.next = curr
		}
		if head == nil {
			head = curr
		}
		prev = curr
	}
	return head
}

func (t *List[T]) String() string {
	var sb strings.Builder
	curr := t
	sb.WriteString("linked list: [")
	for curr != nil {
		fmt.Fprintf(&sb, "val -> %v ", curr.val)
		curr = curr.next
	}
	sb.WriteString("]")

	return sb.String()
}
