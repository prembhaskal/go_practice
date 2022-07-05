package linklist

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	curr := head

	oldToNewMap := make(map[*Node]*Node)

	newPrev := &Node{} // dummy head
	newHead := newPrev

	for curr != nil {
		newCurr := &Node{
			Val:  curr.Val,
			Next: curr.Next,
		}
		newPrev.Next = newCurr // point previous node to newly created node.

		oldToNewMap[curr] = newCurr

		// prepare for next round
		curr = curr.Next
		newPrev = newCurr
	}

	// update random address by navigating both list together
	newCurr := newHead.Next
	curr = head
	for newCurr != nil {
		oldRandom := curr.Random
		newRandom := oldToNewMap[oldRandom]
		newCurr.Random = newRandom

		// next round
		curr = curr.Next
		newCurr = newCurr.Next
	}

	return newHead.Next
}
