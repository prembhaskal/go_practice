package linklist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	// base case
	if head == nil || head.Next == nil {
		return head
	}
	len := 1
	curr := head
	for curr.Next != nil {
		len++
		curr = curr.Next
	}
	// curr  ----  curr.Next
	// 1  2  1
	// 2  3  2
	// 3  4  3
	// 4  5  4
	// 5  nil

	// create cycle, point tail to head
	curr.Next = head

	// fmt.Printf("len is %d\n", len)

	r := k % len    // 2%5 = 2
	r = len - r - 1 // 5-2-1 = 2

	// fmt.Printf("r is %d\n",  r)
	curr = head
	for r > 0 {
		r--
		curr = curr.Next
	}

	newHead := curr.Next
	curr.Next = nil

	return newHead
}
