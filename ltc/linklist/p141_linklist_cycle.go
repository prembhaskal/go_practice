/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 * 
 * https://leetcode.com/problems/linked-list-cycle/
 */
func hasCycle(head *ListNode) bool {
    sptr := head
    fptr := head
    
    hasCycle := false
    
    for ; sptr != nil && fptr != nil ; {
        // move the pointers
        sptr = sptr.Next
        if fptr.Next == nil {
            break
        }
        fptr = fptr.Next.Next
        
        // compare if slow and fast are pointing same.
        if sptr == fptr {
            hasCycle = true
            break
        }
    }
    
    return hasCycle
}