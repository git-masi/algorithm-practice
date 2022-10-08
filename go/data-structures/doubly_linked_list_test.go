package datastructures

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

func TestDoublyLinkedListInsertMethod(t *testing.T) {
	t.Run("the `head` property should be `nil` by default", func(t *testing.T) {
		dll := DoublyLinkedList{}

		if dll.head != nil {
			t.Errorf("expected `nil` got %v", dll.head)
		}
	})

	t.Run("the `tail` property should be `nil` by default", func(t *testing.T) {
		dll := DoublyLinkedList{}

		if dll.tail != nil {
			t.Errorf("expected `nil` got %v", dll.tail)
		}
	})

	t.Run("the `size` property should be `0` by default", func(t *testing.T) {
		dll := DoublyLinkedList{}

		if dll.size != 0 {
			t.Errorf("expected `0` got %v", dll.size)
		}
	})

	t.Run("it should add a new `head` when the first node is inserted", func(t *testing.T) {
		dll := DoublyLinkedList{}
		head := 42
		dll.Insert(42)

		if dll.head.value != head {
			t.Errorf("expected %d got %d", head, dll.head.value)
		}

		if dll.tail.value != head {
			t.Errorf("expected %d got %d", head, dll.tail.value)
		}

		if dll.size != 1 {
			t.Errorf("want 1 got %d", dll.size)
		}
	})

	t.Run("it should add a new `tail` when the second node is inserted", func(t *testing.T) {
		dll := DoublyLinkedList{}
		head := 42
		tail := 9001
		dll.Insert(head)
		dll.Insert(tail)

		if dll.tail.value != tail {
			t.Errorf("expected %d got %d", tail, dll.tail.value)
		}

		if dll.head.next.value != tail {
			t.Errorf("expected %d got %d", tail, dll.head.next.value)
		}

		if dll.tail.prev.value != head {
			t.Errorf("expected %d got %d", head, dll.tail.prev.value)
		}

		if dll.size != 2 {
			t.Errorf("want 2 got %d", dll.size)
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

		if dll.tail.value != tail {
			t.Errorf("expected %d got %d", tail, dll.tail.value)
		}

		if dll.head.next.value != prev {
			t.Errorf("expected %d got %d", prev, dll.head.next.value)
		}

		if dll.tail.prev.value != prev {
			t.Errorf("expected %d got %d", prev, dll.tail.prev.value)
		}

		if dll.size != 3 {
			t.Errorf("want 3 got %d", dll.size)
		}
	})
}

func TestDoublyLinkedListRemoveMethod(t *testing.T) {
	t.Run("It should remove the head node and promote a new head node", func(t *testing.T) {
		dll := DoublyLinkedList{}
		tail := 9001
		dll.Insert(42)
		dll.Insert(tail)
		dll.Remove(dll.head)

		if dll.head.value != tail {
			t.Errorf("expected %d got %d", tail, dll.head.value)
		}
	})

	t.Run("It should remove the tail node", func(t *testing.T) {
		dll := DoublyLinkedList{}
		dll.Insert(42)
		dll.Insert(9001)
		dll.Remove(dll.tail)

		if dll.tail != nil {
			t.Errorf("expected `nil` got %v", dll.tail)
		}
	})

	t.Run("It should update `next` and `prev` values", func(t *testing.T) {
		dll := DoublyLinkedList{}
		dll.Insert(42)
		dll.Insert(9001)
		dll.Insert(1337)
		head := dll.head
		tail := dll.tail
		dll.Remove(head.next)

		if dll.head.next != tail {
			t.Errorf("expected %v got %v", tail, dll.head.next)
		}

		if dll.tail.prev != head {
			t.Errorf("expected %v got %v", head, dll.tail.prev)
		}
	})
}
