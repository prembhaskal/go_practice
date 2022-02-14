package linklist

import (
	"fmt"
	"testing"
)

func TestReverseLinkedList1(t *testing.T) {
	tcs := make([][]int, 0)
	tcs = append(tcs, []int{1, 0, -1, 3, 4, 5, 9})
	tcs = append(tcs, []int{})
	tcs = append(tcs, []int{1})
	tcs = append(tcs, []int{9, 5})
	tcs = append(tcs, []int{9, 5, -1})

	for _, ar := range tcs {
		fmt.Println("new testcase")
		lnklist := createLinkedList(ar)
		printLinkedList(lnklist)

		revlist := reverseList(lnklist)
		printLinkedList(revlist)
	}
}

func printLinkedList(head *ListNode) {
	fmt.Println("********************************")
	curr := head
	for curr != nil {
		fmt.Printf("Val: %d\n", curr.Val)
		curr = curr.Next
	}
}

// create linked list from array
func createLinkedList(list []int) *ListNode {
	var prev, curr *ListNode
	head := &ListNode{}
	prev = head
	for _, v := range list {
		curr = &ListNode{}
		curr.Val = v
		prev.Next = curr
		prev = curr
	}

	return head.Next // return 1st element as head
}
