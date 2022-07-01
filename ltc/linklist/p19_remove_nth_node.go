package linklist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	curr := head

	for n > 0 && curr != nil {
		curr = curr.Next
		n--
	}
	if curr == nil { // it means it is removing from head.
		return head.Next
	}

	// fmt.Printf("curr at : %d\n", curr.Val)

	curr = curr.Next
	prev := head
	for curr != nil {
		curr = curr.Next
		prev = prev.Next
	}

	prev.Next = prev.Next.Next
	return head

}
