package linklist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// return createNewListAndMerge(list1, list2)
	return inplaceMerge(list1, list2)
}

func inplaceMerge(list1 *ListNode, list2 *ListNode) *ListNode {
	curr1 := list1
	curr2 := list2

	tmphead := &ListNode{}
	curr := tmphead
	for curr1 != nil || curr2 != nil {
		if curr1 == nil {
			curr.Next = curr2
			break
		}
		if curr2 == nil {
			curr.Next = curr1
			break
		}

		if curr1.Val <= curr2.Val {
			curr.Next = curr1
			curr1 = curr1.Next
		} else {
			curr.Next = curr2
			curr2 = curr2.Next
		}
		curr = curr.Next
	}

	return tmphead.Next
}

func createNewListAndMerge(list1 *ListNode, list2 *ListNode) *ListNode {
	curr1 := list1
	curr2 := list2

	tmphead := &ListNode{
		Val:  -1,
		Next: nil,
	}
	curr3 := tmphead

	for curr1 != nil || curr2 != nil {
		var newVal int
		if curr1 == nil {
			newVal = curr2.Val
			curr2 = curr2.Next
		} else if curr2 == nil {
			newVal = curr1.Val
			curr1 = curr1.Next
		} else {
			if curr1.Val <= curr2.Val {
				newVal = curr1.Val
				curr1 = curr1.Next
			} else {
				newVal = curr2.Val
				curr2 = curr2.Next
			}
		}

		curr3.Next = &ListNode{Val: newVal}
		curr3 = curr3.Next
	}

	return tmphead.Next
}
