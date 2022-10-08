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
	curr := dll.head

	for {
		if curr == nil {
			return
		}

		if curr == node {
			if curr.prev != nil {
				curr.prev.next = curr.next
			}

			if curr.next != nil {
				curr.next.prev = curr.prev
			}

			if curr == dll.head {
				dll.head = curr.next
			}

			if curr == dll.tail {
				dll.tail = curr.prev
			}

			dll.size--
			return
		}

		curr = curr.next
	}
}

func (dll *DoublyLinkedList) Find(value int) *DoublyLinkedListNode {
	curr := dll.head

	for {
		if curr == nil {
			return nil
		}

		if curr.value == value {
			return curr
		}

		curr = curr.next
	}
}
