package linklist

import (
	"fmt"
	"strings"
)

type MyLinkedList struct {
	head *node // head is 1st element.

}

type node struct {
	val  int
	next *node
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this MyLinkedList) String() string {
	var sb strings.Builder
	curr := this.head

	for i := 0; curr != nil; i++ {
		sb.WriteString(fmt.Sprintf("idx: %d, val: %d  ", i, curr.val))
		curr = curr.next
	}

	return sb.String()
}

func (this *MyLinkedList) Get(index int) int {
	curr := this.head
	for i := 0; i < index; i++ {
		if curr == nil {
			return -1
		}
		curr = curr.next
	}
	if curr == nil {
		return -1
	}

	return curr.val
}

func (this *MyLinkedList) AddAtHead(val int) {
	tmphead := this.head
	head := &node{
		val:  val,
		next: tmphead,
	}
	this.head = head
}

func (this *MyLinkedList) AddAtTail(val int) {
	curr := this.head
	if curr == nil { // empty llist
		this.head = &node{val: val}
		return
	}
	for curr.next != nil { // loop till current tail
		curr = curr.next
	}
	curr.next = &node{val: val}
}

// add a new node, just before node at index - 'index'
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == 0 {
		newNode := &node{
			val: val,
			next: this.head,
		}
		this.head = newNode
		return
	}
	// i == 0 is at head
	
	prev := this.head // prev will finally hold element at i-1 when loop ends
	for i := 1; i < index; i++ {
		if prev == nil {
			break
		}
		curr := prev.next
		prev = curr
	}

	if prev != nil {
		next := prev.next
		prev.next = &node{
			val:  val,
			next: next,
		}
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	curr := this.head
	for i := 1; i < index-1; i++ {
		if curr != nil {
			curr = curr.next
		}
	}

	if curr == nil {
		return
	}

	todelete := curr.next
	if todelete == nil {
		return
	}

	curr.next = todelete.next
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
