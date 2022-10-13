package datastructures

type DoublyLinkedListNode struct {
	next  *DoublyLinkedListNode
	prev  *DoublyLinkedListNode
	value int
}

type DoublyLinkedList struct {
	Head *DoublyLinkedListNode
	Tail *DoublyLinkedListNode
	Size int
}

func (dll *DoublyLinkedList) Insert(value int) *DoublyLinkedListNode {
	node := &DoublyLinkedListNode{value: value}

	if dll.Head == nil {
		dll.Head = node
		dll.Tail = node
	} else {
		dll.Tail.next = node
		node.prev = dll.Tail
		dll.Tail = node
	}

	dll.Size++

	return node
}

func (dll *DoublyLinkedList) Remove(node *DoublyLinkedListNode) {
	if node == nil {
		return
	}

	curr := dll.Head

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

			if curr == dll.Head {
				dll.Head = curr.next
			}

			if curr == dll.Tail {
				dll.Tail = curr.prev
			}

			dll.Size--
			return
		}

		curr = curr.next
	}
}

func (dll *DoublyLinkedList) Find(value int) *DoublyLinkedListNode {
	curr := dll.Head

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
