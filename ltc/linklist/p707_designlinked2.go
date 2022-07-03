package linklist

type MyDLinkedList struct {
	head *DListNode
}

type DListNode struct {
	Prev *DListNode
	Next *DListNode
	Val  int
}

func Constructor2() MyDLinkedList {
	return MyDLinkedList{}
}

func (this *MyDLinkedList) Get(index int) int {
	curr := this.head
	for i := 1; i <= index; i++ {
		if curr == nil {
			break
		}
		curr = curr.Next
	}
	if curr == nil {
		return -1
	}
	return curr.Val
}

func (this *MyDLinkedList) AddAtHead(val int) {
	node := &DListNode{
		Val:  val,
		Next: this.head,
		Prev: nil,
	}

	if this.head != nil {
		this.head.Prev = node
	}

	this.head = node
}

func (this *MyDLinkedList) AddAtTail(val int) {
	if this.head == nil {
		this.AddAtHead(val)
		return
	}
	curr := this.head
	for curr.Next != nil {
		curr = curr.Next
	}
	// curr is at tail now
	node := &DListNode{
		Val:  val,
		Next: nil,
		Prev: curr,
	}

	curr.Next = node
	return
}

func (this *MyDLinkedList) AddAtIndex(index int, val int) {
	if index == 0 {
		this.AddAtHead(val)
		return
	}

	// 0 - 1 - 2 - 3 - 4 - NIL
	// 0 - 1 - NN - 2 - 3 - 4 - NIL
	curr := this.head
	for i := 1; i < index; i++ {
		curr = curr.Next
	}
	if curr == nil {
		return
	}

	node := &DListNode{
		Val:  val,
		Next: curr.Next,
		Prev: curr,
	}

	curr.Next = node
	if node.Next != nil {
		node.Next.Prev = node
	}
	return
}

// 1 - 2  - 3
func (this *MyDLinkedList) DeleteAtIndex(index int) {
	curr := this.head
	for i := 1; i <= index; i++ {
		if curr == nil {
			break
		}
		curr = curr.Next
	}

	if curr == nil {
		return
	}

	prev := curr.Prev
	next := curr.Next
	if prev != nil {
		prev.Next = next
	} else {
		this.head = next
	}

	if next != nil {
		next.Prev = prev
	}
	return
}

/**
 * Your MyDLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
