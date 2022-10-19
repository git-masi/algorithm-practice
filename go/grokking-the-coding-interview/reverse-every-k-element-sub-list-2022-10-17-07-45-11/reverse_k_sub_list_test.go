package main

import (
	"testing"

	linkedlist "github.com/git-masi/algorithm-practice/go/data-structures/linked-list"
)

func TestReverseSubList(t *testing.T) {
	t.Run("It should not reverse the list if k < 2", func(t *testing.T) {
		head := &linkedlist.Node[int]{Value: 0}
		head.Next = &linkedlist.Node[int]{Value: 1}

		got := reverseSubList(head, 1)

		want := &linkedlist.Node[int]{Value: 0}
		want.Next = &linkedlist.Node[int]{Value: 1}

		assertListsAreEqual(t, got, want)
	})

	t.Run("It should reverse the head and the tail", func(t *testing.T) {
		head := &linkedlist.Node[int]{Value: 0}
		head.Next = &linkedlist.Node[int]{Value: 1}

		got := reverseSubList(head, 2)

		want := &linkedlist.Node[int]{Value: 1}
		want.Next = &linkedlist.Node[int]{Value: 0}

		assertListsAreEqual(t, got, want)
	})

	t.Run("It should reverse two sub lists", func(t *testing.T) {
		head := &linkedlist.Node[int]{Value: 1}
		head.Next = &linkedlist.Node[int]{Value: 2}
		head.Next.Next = &linkedlist.Node[int]{Value: 3}
		head.Next.Next.Next = &linkedlist.Node[int]{Value: 4}
		head.Next.Next.Next.Next = &linkedlist.Node[int]{Value: 5}
		head.Next.Next.Next.Next.Next = &linkedlist.Node[int]{Value: 6}

		got := reverseSubList(head, 3)

		want := &linkedlist.Node[int]{Value: 3}
		want.Next = &linkedlist.Node[int]{Value: 2}
		want.Next.Next = &linkedlist.Node[int]{Value: 1}
		want.Next.Next.Next = &linkedlist.Node[int]{Value: 6}
		want.Next.Next.Next.Next = &linkedlist.Node[int]{Value: 5}
		want.Next.Next.Next.Next.Next = &linkedlist.Node[int]{Value: 4}

		assertListsAreEqual(t, got, want)
	})

	t.Run("It should reverse two sub lists and a partial sub list", func(t *testing.T) {
		head := &linkedlist.Node[int]{Value: 1}
		head.Next = &linkedlist.Node[int]{Value: 2}
		head.Next.Next = &linkedlist.Node[int]{Value: 3}
		head.Next.Next.Next = &linkedlist.Node[int]{Value: 4}
		head.Next.Next.Next.Next = &linkedlist.Node[int]{Value: 5}
		head.Next.Next.Next.Next.Next = &linkedlist.Node[int]{Value: 6}
		head.Next.Next.Next.Next.Next.Next = &linkedlist.Node[int]{Value: 7}
		head.Next.Next.Next.Next.Next.Next.Next = &linkedlist.Node[int]{Value: 8}

		got := reverseSubList(head, 3)

		want := &linkedlist.Node[int]{Value: 3}
		want.Next = &linkedlist.Node[int]{Value: 2}
		want.Next.Next = &linkedlist.Node[int]{Value: 1}
		want.Next.Next.Next = &linkedlist.Node[int]{Value: 6}
		want.Next.Next.Next.Next = &linkedlist.Node[int]{Value: 5}
		want.Next.Next.Next.Next.Next = &linkedlist.Node[int]{Value: 4}
		want.Next.Next.Next.Next.Next.Next = &linkedlist.Node[int]{Value: 8}
		want.Next.Next.Next.Next.Next.Next.Next = &linkedlist.Node[int]{Value: 7}

		assertListsAreEqual(t, got, want)
	})
}

// There is probably a better way to display feedback. Ideally I would see all of the
// nodes leading up to the point of failure
func assertListsAreEqual(t *testing.T, got, want *linkedlist.Node[int]) {
	for {
		if got == nil && want == nil {
			break
		}

		if got == nil && want != nil {
			t.Errorf("got %v, want %d", got, want.Value)
			break
		}

		if got != nil && want == nil {
			t.Errorf("got %d, want %v", got.Value, want)
			break
		}

		if got.Value != want.Value {
			t.Errorf("got %d, want %d", got.Value, want.Value)
			break
		}

		got = got.Next
		want = want.Next
	}
}
