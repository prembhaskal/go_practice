package linklist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	seen := make(map[*ListNode]bool, 0)

	curr := head

	var cycleNode *ListNode
	for curr != nil {
		// check if already seen
		if seen[curr] {
			cycleNode = curr
			break
		}

		seen[curr] = true
		// move to next
		curr = curr.Next
	}

	return cycleNode
}
