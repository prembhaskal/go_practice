package recurse

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

func swapPairs(head *ListNode) *ListNode {
	// return swapPairsRec(head)
	return swapiter(head)
}
func swapPairsRec(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	tmp := head
	next := head.Next
	if next == nil {
		return head
	}

	nextHead := swapPairsRec(next.Next)
	next.Next = tmp
	tmp.Next = nextHead
	return next
}

// TODO - optimize
func swapiter(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	curr := head
	newHead := curr.Next
	var prev *ListNode
	for curr != nil && curr.Next != nil {
		tmp := curr
		next := curr.Next
		tmp.Next = next.Next
		next.Next = tmp

		if prev != nil {
			prev.Next = next
		}

		// next round
		curr = tmp.Next
		prev = tmp
	}

	return newHead
}
