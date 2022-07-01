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
    if head == nil {
        return head
    }
    
    len := 1
    curr := head
    for curr.Next != nil {
        len++
        curr = curr.Next
    }
  
    // create cycle, point tail to head
    curr.Next = head    
    r := k % len // 2%5 = 2
    r = len - r - 1  // 5-2-1 = 2
    curr = head
    for ;r > 0;r--  {
        curr = curr.Next
    }
    
    newHead := curr.Next
    curr.Next = nil
    
    return newHead
}