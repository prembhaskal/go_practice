package linklist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	tmphead := &ListNode{Next: head}

	prev := tmphead
	curr := head
	for curr != nil {
		if curr.Val == val {
			prev.Next = curr.Next
			curr = curr.Next
		} else {
			prev = curr
			curr = curr.Next
		}
	}

	return tmphead.Next
}
