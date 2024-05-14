package linklist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 // maintain two pointer, forward pointer (ahead n steps) and lagging pointer
 // when forward pointer reaches end, lagging pointer is at right place
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	curr := head

	// curr is the forward pointer
	for n > 0 && curr != nil {
		curr = curr.Next
		n--
	}
	if curr == nil { // it means it is removing from head.
		return head.Next
	}

	// fmt.Printf("curr at : %d\n", curr.Val)

	curr = curr.Next
	prev := head // prev is lagging pointer
	for curr != nil {
		curr = curr.Next
		prev = prev.Next
	}
	// when forward pointer is at end, it means lagging pointer is at right place.

	prev.Next = prev.Next.Next
	return head
}
