package linkedlist

import (
	"reflect"
	"testing"
)

func TestLinkedListNode(t *testing.T) {
	t.Run("Properties should have zero values by default", func(t *testing.T) {
		node := Node[int]{}

		assertNodeValue(t, node, 0)

		if node.Next != nil {
			t.Error("The next node should be 'nil' by default")
		}
	})
}

func TestLinkedList(t *testing.T) {
	t.Run("It should have zero values by default", func(t *testing.T) {
		l := LinkedList[int]{}
		size := 0

		if l.Head != nil {
			t.Error("The head node should be 'nil' by default")
		}

		if l.Tail != nil {
			t.Error("The tail node should be 'nil' by default")
		}

		assertSize(t, l, size)
	})
}

func TestLinkedListInsert(t *testing.T) {
	t.Run("It should add a new `head` when the first node is inserted", func(t *testing.T) {
		l := LinkedList[int]{}
		head := 42
		size := 1
		l.Insert(head)

		assertNodeValue(t, *l.Head, head)

		assertNodeValue(t, *l.Tail, head)

		assertSize(t, l, size)
	})

	t.Run("It should add a new `tail` when the second node is inserted", func(t *testing.T) {
		l := LinkedList[int]{}
		head := 42
		tail := 1337
		size := 2
		l.Insert(head)
		l.Insert(tail)

		assertNodeValue(t, *l.Head, head)

		assertNodeValue(t, *l.Tail, tail)

		assertSize(t, l, size)
	})

	t.Run("All nodes should be reachable from the head", func(t *testing.T) {
		l := LinkedList[int]{}
		head := 42
		next := 9001
		tail := 1337
		size := 3
		l.Insert(head)
		l.Insert(next)
		l.Insert(tail)

		assertNodeValue(t, *l.Head, head)

		assertNodeValue(t, *l.Head.Next, next)

		assertNodeValue(t, *l.Head.Next.Next, tail)

		assertNodeValue(t, *l.Tail, tail)

		assertSize(t, l, size)
	})
}

func TestLinkedListToSlice(t *testing.T) {
	t.Run("It should return a slice of linked list values", func(t *testing.T) {
		want := []int{1, 2, 3, 4}
		ll := LinkedList[int]{}

		for _, v := range want {
			ll.Insert(v)
		}

		got := ll.ToSlice()

		if !reflect.DeepEqual(want, got) {
			t.Errorf("want %v, got %v ", want, got)
		}
	})
}

func TestLinkedListRemoveHead(t *testing.T) {
	t.Run("It should not remove the head of an empty list", func(t *testing.T) {
		ll := LinkedList[int]{}

		ll.RemoveHead()

		if ll.Head != nil {
			t.Error("An empty list should be unchanged")
		}
	})

	t.Run("I should remove the head and the list should be empty", func(t *testing.T) {
		ll := LinkedList[int]{}

		ll.Insert(42)

		want := ll.Head
		got := ll.RemoveHead()

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}

		if ll.Head != nil || ll.Tail != nil {
			t.Error("Expected head and tail to both be 'nil'")
		}
	})

	t.Run("It should remove the head and the new head and tail should be equal", func(t *testing.T) {
		ll := LinkedList[int]{}
		tail := 1337

		ll.Insert(42)
		ll.Insert(tail)

		want := ll.Head
		got := ll.RemoveHead()

		if got.Next != nil {
			t.Errorf("Expected the Next value to be 'nil'")
		}

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}

		assertNodeValue(t, *ll.Head, tail)

		assertNodeValue(t, *ll.Tail, tail)
	})

	t.Run("It should remove the head and promote a new head", func(t *testing.T) {
		ll := LinkedList[int]{}
		head := 42
		next := 9001
		tail := 1337

		ll.Insert(head)
		ll.Insert(next)
		ll.Insert(tail)

		want := ll.Head
		got := ll.RemoveHead()

		if got.Next != nil {
			t.Errorf("Expected the Next value to be 'nil'")
		}

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}

		assertNodeValue(t, *ll.Head, next)

		assertNodeValue(t, *ll.Tail, tail)
	})
}

func assertNodeValue(t *testing.T, n Node[int], v int) {
	if n.Value != v {
		t.Errorf("got %d, want %d", n.Value, v)
	}
}

func assertSize(t *testing.T, l LinkedList[int], s int) {
	if l.Size != s {
		t.Errorf("got %d, want %d", l.Size, s)
	}
}
