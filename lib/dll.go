package lib

type Node[K any] struct {
	data K
	prev *Node[K]
	next *Node[K]
}

type DoublyLinkedList[K any] struct {
	len  int
	Head *Node[K]
	Tail *Node[K]
}

func NewNode[K any](data K) *Node[K] {
	return &Node[K]{data: data}
}

func (node *Node[K]) GetKey() K {
	return node.data
}
func NewtDoublyList[K any]() *DoublyLinkedList[K] {
	return &DoublyLinkedList[K]{}
}

func (d *DoublyLinkedList[K]) AddEndNodeDLL(newNode *Node[K]) {

	if d.Head == nil {
		d.Head = newNode
		d.Tail = newNode
	} else {
		newNode.prev = d.Tail
		d.Tail.next = newNode
		d.Tail = newNode
	}
	d.len++
	return
}

func (d *DoublyLinkedList[K]) Size() int {
	return d.len
}

func (d *DoublyLinkedList[K]) DetachNode(node *Node[K]) {
	// Just Simply modifying the pointers.

	prev := node.prev
	next := node.next

	if prev != nil {
		prev.next = next
	}
	if next != nil {
		next.prev = prev
	}

	node.prev = nil
	node.next = nil

	if node == d.Head {

		d.Head = next
	}

	if node == d.Tail {

		d.Tail = prev
	}

	d.len--
	return
}
