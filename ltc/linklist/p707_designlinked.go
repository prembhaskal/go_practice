package linklist

import (
	"fmt"
	"strings"
)

// https://leetcode.com/problems/design-linked-list/

// TODO try using a dummy head , sentinel pointer for simplifying loop, write separate impl.
// TODO try keeping track of length and tail, to reduce O() complexity.
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

	sb.WriteString("HEAD --> ")
	for i := 0; curr != nil; i++ {
		sb.WriteString(fmt.Sprintf("[ i: %d, v: %d ]", i, curr.val))
		sb.WriteString(" --> ")
		curr = curr.next
	}
	sb.WriteString("NIL")

	return sb.String()
}

func (this *MyLinkedList) Get(idx int) int {
	curr := this.head
	i := 0
	for curr != nil {
		if i == idx {
			return curr.val
		}
		i++
		curr = curr.next
	}

	return -1

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
	if curr == nil { // empty list
		this.head = &node{val: val}
		return
	}
	for curr.next != nil { // loop till current tail
		curr = curr.next
	}
	curr.next = &node{val: val}
}

// add a new node so that new node gets index 'index'
func (this *MyLinkedList) AddAtIndex(idx int, val int) {
	if idx == 0 {
		this.AddAtHead(val)
		return
	}
	// i == 0 is at head
	// find element at (idx-1)th index.
	prev := this.head // prev will finally hold element at i-1 when loop ends
	i := 0
	for prev != nil {
		// here prev will point to node at index i
		if i == idx-1 {
			break
		}
		i++
		prev = prev.next
	}

	if prev != nil {
		next := prev.next
		prev.next = &node{
			val:  val,
			next: next,
		}
	}
}

func (this *MyLinkedList) DeleteAtIndex(idx int) {
	// delete at head, for idx = 0
	if idx == 0 {
		prev := this.head
		if prev != nil {
			this.head = prev.next
		}
		return
	}

	// find node at idx-1 index
	prev := this.head
	i := 0
	for prev != nil {
		if i == idx-1 {
			break
		}
		i++
		prev = prev.next
	}

	// prev -> curr -> next
	// point prev to next
	if prev != nil {
		curr := prev.next
		if curr == nil {
			return
		}
		next := curr.next
		prev.next = next
	}
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
