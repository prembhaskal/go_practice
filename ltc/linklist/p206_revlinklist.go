package linklist

import (
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// prev -> curr -> next  ==> prev <- curr <- next
// here head is the 1st element (with a value)
// https://leetcode.com/problems/reverse-linked-list/
func reverseList(head *ListNode) *ListNode {
	var curr, prev, next *ListNode

	prev = nil
	curr = head
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

func reverseBy3ptrs(head *ListNode) *ListNode {
	// prev ->  curr  -> next -> next1 -> next2 -> next3 -> NIL
	// prev1 <- prev     curr -> next -> next2 -> next3 -> NIL
	// prev2 <- prev1 <- prev    next1 -> next2 -> next3 -> NIL
	// prev <-  curr  <- next ...         prev <-  curr -> NIL

	var prev, curr *ListNode
	prev = nil
	curr = head

	for curr != nil {
		// fmt.Printf("curr is %d\n", curr.Val)
		// printList(prev)

		next := curr.Next
		curr.Next = prev

		prev = curr
		curr = next
	}

	return prev
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%2d", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func reverseByPushingOnTop(head *ListNode) *ListNode {
	// way one, remove node and put on head
	if head == nil {
		return nil
	}
	curr := head
	tmpHead := head
	for curr.Next != nil {
		next := curr.Next
		curr.Next = curr.Next.Next

		// change head
		next.Next = tmpHead
		tmpHead = next
	}

	return tmpHead
}
