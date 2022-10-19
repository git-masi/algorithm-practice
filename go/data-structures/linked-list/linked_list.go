package linkedlist

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type LinkedList[T any] struct {
	Head *Node[T]
	Tail *Node[T]
	Size int
}

func (ll *LinkedList[T]) Insert(v T) {
	node := &Node[T]{Value: v}

	if ll.Head == nil {
		ll.Head = node
		ll.Tail = node
	} else {
		ll.Tail.Next = node
		ll.Tail = node
	}

	ll.Size++
}

func (ll *LinkedList[T]) RemoveHead() *Node[T] {
	if ll.Head == nil {
		return nil
	}

	head := ll.Head
	next := head.Next

	if head == ll.Tail {
		ll.Tail = nil
	}

	head.Next = nil
	ll.Head = next

	return head
}

func (ll LinkedList[T]) ToSlice() []T {
	values := []T{}

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

func CreateTestList(size int) LinkedList[int] {
	ll := LinkedList[int]{}

	if size > 0 {
		for i := 0; i < size; i++ {
			ll.Insert(i)
		}
	} else {
		for i := (size * -1) - 1; i >= 0; i-- {
			ll.Insert(i)
		}
	}

	return ll
}
