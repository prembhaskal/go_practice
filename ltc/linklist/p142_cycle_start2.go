package linklist
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
    sptr := head
    fptr := head

    hasCycle := false
    for ; sptr != nil && fptr != nil ; {
        sptr = sptr.Next
        if fptr.Next == nil {
            break
        }
        fptr = fptr.Next.Next

        // check
        if sptr == fptr {
            hasCycle = true
            break
        }
    }

    if !hasCycle {
        return nil
    }

    // start a pointer from start and continue the sptr too. check where they meet.
    nptr := head    
    for fptr != nptr {
        fptr = fptr.Next
        nptr = nptr.Next
    }
    return nptr
}