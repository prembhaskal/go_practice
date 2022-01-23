package revlinklist

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
