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
		dll.tail = node
	} else if dll.tail == nil {
		dll.head.next = node
		dll.tail = node
		node.prev = dll.head
	} else {
		dll.tail.next = node
		node.prev = dll.tail
		dll.tail = node
	}

	dll.size++
}

func (dll *DoublyLinkedList) Remove(node *DoublyLinkedListNode) {
	if node == dll.head {
		dll.head = dll.head.next
		dll.size--
	} else if node == dll.tail {
		dll.tail.prev.next = nil
		dll.tail = nil
		dll.size--
	} else {
		curr := dll.head

		for {
			if curr == nil {
				return
			}

			if curr == node {
				curr.prev.next = curr.next
				curr.next.prev = curr.prev
				dll.size--
				return
			}

			curr = curr.next
		}
	}
}
