package recurse

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func mergeKLists(lists []*ListNode) *ListNode {
    return mergeR(0, len(lists)-1, lists)
}

// do a merge sort kind.
func mergeR(start, end int, lists []*ListNode) *ListNode {
    if start > end {
        return nil
    }
    if start == end {
        return lists[start]
    }

    mid := (start + end) / 2

    A := mergeR(start, mid, lists)
    B := mergeR(mid+1, end, lists)
    return merge(A, B)
}

func merge(a, b *ListNode) *ListNode {
    c := &ListNode{
        Val: 0,
        Next: nil,
    }
    // dummy head
    head := c

    for a != nil && b != nil {
        if a.Val < b.Val {
            c.Next = a
            a = a.Next
        } else {
            c.Next = b
            b = b.Next
        } 
        c = c.Next
    }
    if a != nil {
        c.Next = a
    } else {
        c.Next = b
    }
    
    return head.Next // ignore dummy head
}