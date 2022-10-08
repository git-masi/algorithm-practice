package datastructures

type DoublyLinkedListNode struct {
	next  *DoublyLinkedListNode
	prev  *DoublyLinkedListNode
	value int
}

type DoublyLinkedList struct {
	head *DoublyLinkedListNode
	tail *DoublyLinkedListNode
	size int
}

func (dll *DoublyLinkedList) Insert(value int) {
	node := &DoublyLinkedListNode{value: value}

	if dll.head == nil {
		dll.head = node
	} else if dll.tail == nil {
		dll.head.next = node
		dll.tail = node
		node.prev = dll.head
	} else {
		dll.tail.next = node
		node.prev = dll.tail
		dll.tail = node
	}
}
