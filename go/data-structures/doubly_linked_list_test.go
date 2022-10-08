package datastructures

import "testing"

func TestDoublyLinkedListNode(t *testing.T) {
	t.Run("the `next` property should be `nil` by default", func(t *testing.T) {
		dlln := DoublyLinkedListNode{}

		if dlln.next != nil {
			t.Errorf("expected `nil` got %v", dlln.next)
		}
	})

	t.Run("the `prev` property should be `nil` by default", func(t *testing.T) {
		dlln := DoublyLinkedListNode{}

		if dlln.prev != nil {
			t.Errorf("expected `nil` got %v", dlln.prev)
		}
	})

	t.Run("the `value` property should be `0` by default", func(t *testing.T) {
		dlln := DoublyLinkedListNode{}

		if dlln.value != 0 {
			t.Errorf("expected `nil` got %v", dlln.value)
		}
	})
}
