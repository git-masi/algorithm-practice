package linkedlist

type LinkedListNode struct {
	Value int
	Next  *LinkedListNode
}

type LinkedList struct {
	Head *LinkedListNode
	Tail *LinkedListNode
	Size int
}

func (l *LinkedList) Insert(v int) {
	node := &LinkedListNode{Value: v}

	if l.Head == nil {
		l.Head = node
		l.Tail = node
	} else {
		l.Tail.Next = node
		l.Tail = node
	}

	l.Size++
}
