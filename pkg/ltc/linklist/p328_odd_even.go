package linklist

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	odd := head
	even := head.Next
	firstEven := even
	for odd != nil && odd.Next != nil && even != nil && even.Next != nil {
		fmt.Printf("odd: %d, even: %d\n", odd.Val, even.Val)

		newOdd := even.Next
		newEven := even.Next.Next

		odd.Next = newOdd
		even.Next = newEven

		odd = newOdd
		even = newEven
	}
	if odd != nil {
		fmt.Printf("odd: %d\n", odd.Val)
	}
	if even != nil {
		fmt.Printf("even: %d\n", even.Val)
	}

	odd.Next = firstEven

	return head
}
