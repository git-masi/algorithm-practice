package linkedlist

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

func (ll *LinkedList) Insert(v int) {
	node := &Node{Value: v}

	if ll.Head == nil {
		ll.Head = node
		ll.Tail = node
	} else {
		ll.Tail.Next = node
		ll.Tail = node
	}

	ll.Size++
}

func (ll *LinkedList) RemoveHead() *Node {
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

func CreateTestList(size int) LinkedList {
	ll := LinkedList{}

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
