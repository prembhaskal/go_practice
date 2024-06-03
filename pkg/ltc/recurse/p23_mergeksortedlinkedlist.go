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


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 // Solved using min heap
 // add top from all lists to minheap
 // extract from minheap and append to list, add next element of top from same list back into heap (if present)
 // not adding everything in heap, to keep heap size at most 'k'
 // runtime O(n*log(k))
 func mergeKLists1(lists []*ListNode) *ListNode {
    mh := newminheap()

    // add top elements from all
    for i := 0; i < len(lists); i++ {
        top := lists[i]
        if top != nil {
            mh.add(&node{Val: top.Val, idx: i, actnode: top})
        }
    }

    newlist := &ListNode {
        Val: -1,
        Next: nil,
    }
    dummyhead := newlist

    for mh.size > 0 {
        minval := mh.extractMin()
        actnode := minval.actnode

        newlist.Next = actnode
        newlist = newlist.Next

        if actnode.Next != nil {
            next := actnode.Next
            mh.add(&node{Val: next.Val, idx: minval.idx, actnode: next})
        }

    }

    return dummyhead.Next
    
}

// heap impl
type node struct {
    Val int
    idx int
    actnode *ListNode
}

type minheap struct {
    ar []*node
    size int
}

func newminheap() *minheap {
    ar := make([]*node, 1)
    ar[0] = nil // dummy entry at 0, actual starts at 1
    return &minheap{
        ar: ar,
        size: 0,
    }
}

func (m *minheap) parent(i int) int {
    return i / 2
}

func (m *minheap) left(i int) int {
    return 2 * i
}
func (m * minheap) right(i int) int {
    return 2 * i + 1
}

func (m *minheap) fixdown(i int) {
    L := m.left(i)
    R := m.right(i)

    smallest := i
    if L <= m.size && m.ar[L].Val < m.ar[smallest].Val {
        smallest = L
    }
    if R <= m.size && m.ar[R].Val < m.ar[smallest].Val {
        smallest = R
    }
    if i != smallest {
        // swap
        m.ar[i], m.ar[smallest] = m.ar[smallest], m.ar[i]
        m.fixdown(smallest)
    }
}

func (m *minheap) fixup(i int) {
    par := m.parent(i)
    for par != 0 && m.ar[par].Val > m.ar[i].Val {
        m.ar[par], m.ar[i] = m.ar[i], m.ar[par]
        i = par
        par = m.parent(i)
    }
}

func (m *minheap) add(item *node) {
    m.ar = append(m.ar, item)
    m.size++
    m.fixup(m.size)
}

func (m *minheap) extractMin() *node {
    if m.size == 0 {
        panic("empty heap")
    }
    minval := m.ar[1]
    m.ar[1] = m.ar[m.size]
    m.ar = m.ar[:m.size]
    m.size--

    m.fixdown(1)
    return minval
}
