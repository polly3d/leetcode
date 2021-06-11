package ds

type MyLinkedList struct {
	size   int
	header *Node
	tailer *Node
}

type Node struct {
	val  int
	pre  *Node
	succ *Node
}

/** Initialize your data structure here. */
func NewLinkList() MyLinkedList {
	header := &Node{val: -1}
	tailer := &Node{val: -1}
	header.succ = tailer
	tailer.pre = header
	return MyLinkedList{size: 0, header: header, tailer: tailer}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}
	node := this.header.succ
	for index > 0 {
		index--
		node = node.succ
	}
	return node.val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	first := this.header.succ

	n := &Node{val: val, pre: this.header, succ: first}
	this.header.succ = n
	first.pre = n

	this.size++

}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	last := this.tailer.pre
	n := &Node{val: val, pre: last, succ: this.tailer}
	this.tailer.pre = n
	last.succ = n
	this.size++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.size {
		return
	}
	n := this.header.succ
	for index > 0 {
		index--
		n = n.succ
	}
	newNode := &Node{val: val, pre: n.pre, succ: n}

	n.pre.succ = newNode
	n.pre = newNode
	this.size++
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index > this.size-1 {
		return
	}
	node := this.header.succ
	for index > 0 {
		index--
		node = node.succ
	}
	node.pre.succ = node.succ
	node.succ.pre = node.pre
	this.size--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
