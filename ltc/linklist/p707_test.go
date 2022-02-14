package linklist

import (
	"fmt"
	"testing"
)

func TestLinkedListDesign(t *testing.T) {
	// * Your MyLinkedList object will be instantiated and called as such:
	obj := Constructor()
	param1 := obj.Get(1)
	fmt.Printf("val: %d\n", param1)
	obj.AddAtHead(3)
	// obj.AddAtHead(2)
	// obj.AddAtHead(1)
	fmt.Printf("linkedlist: %s\n", obj)
	// obj.AddAtTail(4)
	// obj.AddAtTail(5)
	// obj.AddAtTail(6)
	// fmt.Printf("linkedlist: %s\n", obj)
	obj.AddAtIndex(0, 7)
	fmt.Printf("add '%d' at index '%d'\n", 7, 0)
	fmt.Printf("linkedlist: %s\n\n", obj)
	// param1 = obj.Get(1)
	// fmt.Printf("val: %d\n", param1)
	// obj.AddAtTail(3)
	// fmt.Printf("linkedlist: %s\n", obj)
	obj.AddAtIndex(1,2)
	fmt.Printf("add '%d' at index '%d'\n", 2, 1)
	fmt.Printf("linkedlist: %s\n\n", obj)
	
	obj.AddAtIndex(2,5)
	fmt.Printf("add '%d' at index '%d'\n", 5, 2)
	fmt.Printf("linkedlist: %s\n\n", obj)
	
	obj.AddAtIndex(9,10)
	fmt.Printf("add '%d' at index '%d'\n", 10, 9)
	fmt.Printf("linkedlist: %s\n\n", obj)

	fmt.Printf("value at index:%d, val: %d\n", 0, obj.Get(0))
	fmt.Printf("value at index:%d, val: %d\n", 1, obj.Get(1))
	fmt.Printf("value at index:%d, val: %d\n", 2, obj.Get(2))
	fmt.Printf("value at index:%d, val: %d\n", 3, obj.Get(3))
	fmt.Printf("value at index:%d, val: %d\n", 5, obj.Get(5))
	
	obj.DeleteAtIndex(2)
	fmt.Printf("delete at index: '%d'\n", 2)
	fmt.Printf("linkedlist: %s\n\n", obj)


	obj.DeleteAtIndex(0)
	fmt.Printf("delete at index '%d'\n", 0)
	fmt.Printf("linkedlist: %s\n\n", obj)
	
	obj.DeleteAtIndex(3)
	fmt.Printf("delete at index '%d'\n", 3)
	fmt.Printf("linkedlist: %s\n\n", obj)
}