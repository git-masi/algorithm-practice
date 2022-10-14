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

func (ll *LinkedList) Insert(v int) {
	node := &LinkedListNode{Value: v}

	if ll.Head == nil {
		ll.Head = node
		ll.Tail = node
	} else {
		ll.Tail.Next = node
		ll.Tail = node
	}

	ll.Size++
}

func (ll LinkedList) ToSlice() []int {
	values := []int{}

	curr := ll.Head

	for {
		if curr == nil {
			break
		}

		values = append(values, curr.Value)

		curr = curr.Next
	}

	return values
}
