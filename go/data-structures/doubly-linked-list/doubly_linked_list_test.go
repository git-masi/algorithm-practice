package doublylinkedlist

import "testing"

func TestDoublyLinkedListNode(t *testing.T) {
	t.Run("the `next` property should be `nil` by default", func(t *testing.T) {
		node := DoublyLinkedListNode{}

		if node.next != nil {
			t.Errorf("expected `nil` got %v", node.next)
		}
	})

	t.Run("the `prev` property should be `nil` by default", func(t *testing.T) {
		node := DoublyLinkedListNode{}

		if node.prev != nil {
			t.Errorf("expected `nil` got %v", node.prev)
		}
	})

	t.Run("the `value` property should be `0` by default", func(t *testing.T) {
		node := DoublyLinkedListNode{}

		if node.value != 0 {
			t.Errorf("expected `0` got %v", node.value)
		}
	})
}

func TestDoublyLinkedListInsert(t *testing.T) {
	t.Run("Properties should have zero values for their type by default", func(t *testing.T) {
		dll := DoublyLinkedList{}

		if dll.Head != nil {
			t.Errorf("want `nil` got %v", dll.Head)
		}

		if dll.Tail != nil {
			t.Errorf("want `nil` got %v", dll.Tail)
		}

		if dll.Size != 0 {
			t.Errorf("want `0` got %v", dll.Size)
		}
	})

	t.Run("it should add a new `head` when the first node is inserted", func(t *testing.T) {
		dll := DoublyLinkedList{}
		head := 42
		dll.Insert(42)

		if dll.Head.value != head {
			t.Errorf("expected %d got %d", head, dll.Head.value)
		}

		if dll.Tail.value != head {
			t.Errorf("expected %d got %d", head, dll.Tail.value)
		}

		if dll.Size != 1 {
			t.Errorf("want 1 got %d", dll.Size)
		}
	})

	t.Run("it should add a new `tail` when the second node is inserted", func(t *testing.T) {
		dll := DoublyLinkedList{}
		head := 42
		tail := 9001
		dll.Insert(head)
		dll.Insert(tail)

		if dll.Tail.value != tail {
			t.Errorf("expected %d got %d", tail, dll.Tail.value)
		}

		if dll.Head.next.value != tail {
			t.Errorf("expected %d got %d", tail, dll.Head.next.value)
		}

		if dll.Tail.prev.value != head {
			t.Errorf("expected %d got %d", head, dll.Tail.prev.value)
		}

		if dll.Size != 2 {
			t.Errorf("want 2 got %d", dll.Size)
		}
	})

	t.Run("it should update the tail when the third node is added", func(t *testing.T) {
		dll := DoublyLinkedList{}
		head := 42
		prev := 9001
		tail := 1377
		dll.Insert(head)
		dll.Insert(prev)
		dll.Insert(tail)

		if dll.Tail.value != tail {
			t.Errorf("expected %d got %d", tail, dll.Tail.value)
		}

		if dll.Head.next.value != prev {
			t.Errorf("expected %d got %d", prev, dll.Head.next.value)
		}

		if dll.Tail.prev.value != prev {
			t.Errorf("expected %d got %d", prev, dll.Tail.prev.value)
		}

		if dll.Size != 3 {
			t.Errorf("want 3 got %d", dll.Size)
		}
	})

	t.Run("It should return the inserted node", func(t *testing.T) {
		dll := DoublyLinkedList{}
		tail := 8
		dll.Insert(42)
		dll.Insert(9001)
		dll.Insert(1337)
		node := dll.Insert(tail)

		if node.value != tail {
			t.Errorf("expected %d got %d", tail, node.value)
		}
	})
}

func TestDoublyLinkedListRemove(t *testing.T) {
	t.Run("It should remove the current head node and promote a new head node", func(t *testing.T) {
		dll := DoublyLinkedList{}
		tail := 9001
		dll.Insert(42)
		dll.Insert(tail)
		dll.Remove(dll.Head)

		if dll.Head.value != tail {
			t.Errorf("expected %d got %d", tail, dll.Head.value)
		}

		if dll.Head.prev != nil {
			t.Errorf("want `nil` got %v", dll.Head.prev)
		}

		if dll.Size != 1 {
			t.Errorf("want 1 got %d", dll.Size)
		}
	})

	t.Run("It should remove the tail node and promote a new tail node", func(t *testing.T) {
		dll := DoublyLinkedList{}
		dll.Insert(42)
		dll.Insert(9001)
		dll.Remove(dll.Tail)

		if dll.Tail != dll.Head {
			t.Errorf("want %v got %v", dll.Head, dll.Tail)
		}

		if dll.Tail.next != nil {
			t.Errorf("want `nil` got %v", dll.Tail.next)
		}

		if dll.Size != 1 {
			t.Errorf("want 1 got %d", dll.Size)
		}
	})

	t.Run("It should update `next` and `prev` values", func(t *testing.T) {
		dll := DoublyLinkedList{}
		dll.Insert(42)
		dll.Insert(9001)
		dll.Insert(1337)
		head := dll.Head
		tail := dll.Tail
		dll.Remove(head.next)

		if dll.Head.next != tail {
			t.Errorf("expected %v got %v", tail, dll.Head.next)
		}

		if dll.Tail.prev != head {
			t.Errorf("expected %v got %v", head, dll.Tail.prev)
		}

		if dll.Size != 2 {
			t.Errorf("want 2 got %d", dll.Size)
		}
	})

	t.Run("It should do nothing if attempting to remove a `nil` node", func(t *testing.T) {
		dll := DoublyLinkedList{}
		head := 42
		dll.Insert(head)
		dll.Remove(nil)

		if dll.Head.value != head {
			t.Errorf("want %v got %v", head, dll.Head.value)
		}

		if dll.Tail.value != head {
			t.Errorf("want %v got %v", head, dll.Tail.value)
		}
	})
}

func TestDoublyLinkedListFind(t *testing.T) {
	t.Run("It should return `nil` if there are no nodes in the list", func(t *testing.T) {
		dll := DoublyLinkedList{}
		got := dll.Find(42)

		if got != nil {
			t.Errorf("got %v want `nil`", got)
		}
	})

	t.Run("It should return `nil` if a node with the specified value can't be found", func(t *testing.T) {
		dll := DoublyLinkedList{}
		dll.Insert(42)
		got := dll.Find(1337)

		if got != nil {
			t.Errorf("got %v want `nil`", got)
		}
	})

	t.Run("It should return the node with the specified value", func(t *testing.T) {
		dll := DoublyLinkedList{}
		dll.Insert(42)
		dll.Insert(9001)
		dll.Insert(1337)
		got := dll.Find(9001)

		if got != dll.Head.next {
			t.Errorf("got %v want %v", got, dll.Head.next)
		}
	})
}
