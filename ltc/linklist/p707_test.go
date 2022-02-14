package linklist

import (
	"fmt"
	"testing"
)

func TestLinkedList(t *testing.T) {
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
	fmt.Printf("linkedlist: %s\n", obj)
	// param1 = obj.Get(1)
	// fmt.Printf("val: %d\n", param1)
	// obj.AddAtTail(3)
	// fmt.Printf("linkedlist: %s\n", obj)
	obj.AddAtIndex(1,2)
	fmt.Printf("linkedlist: %s\n", obj)
	obj.AddAtIndex(10,10)
	fmt.Printf("linkedlist: %s\n", obj)
	obj.DeleteAtIndex(0)
	fmt.Printf("linkedlist: %s\n", obj)
}