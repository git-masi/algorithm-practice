package linkedlist

import (
	"reflect"
	"testing"
)

func TestLinkedListNode(t *testing.T) {
	t.Run("Properties should have zero values by default", func(t *testing.T) {
		node := LinkedListNode{}

		assertNodeValue(t, node, 0)

		if node.Next != nil {
			t.Error("The next node should be 'nil' by default")
		}
	})
}

func TestLinkedList(t *testing.T) {
	t.Run("It should have zero values by default", func(t *testing.T) {
		l := LinkedList{}
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
		l := LinkedList{}
		head := 42
		size := 1
		l.Insert(head)

		assertNodeValue(t, *l.Head, head)

		assertNodeValue(t, *l.Tail, head)

		assertSize(t, l, size)
	})

	t.Run("It should add a new `tail` when the second node is inserted", func(t *testing.T) {
		l := LinkedList{}
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
		l := LinkedList{}
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
		ll := LinkedList{}

		for _, v := range want {
			ll.Insert(v)
		}

		got := ll.ToSlice()

		if !reflect.DeepEqual(want, got) {
			t.Errorf("want %v, got %v ", want, got)
		}
	})
}

func assertNodeValue(t *testing.T, n LinkedListNode, v int) {
	if n.Value != v {
		t.Errorf("got %d, want %d", n.Value, v)
	}
}

func assertSize(t *testing.T, l LinkedList, s int) {
	if l.Size != s {
		t.Errorf("got %d, want %d", l.Size, s)
	}
}
