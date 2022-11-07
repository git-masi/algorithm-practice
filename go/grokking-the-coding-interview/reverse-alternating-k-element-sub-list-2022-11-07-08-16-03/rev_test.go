package main

import (
	"reflect"
	"testing"

	list "github.com/git-masi/algorithm-practice/go/data-structures/linked-list"
)

func TestRev(t *testing.T) {
	t.Run("it should not mutate the list if there is only 1 node", func(t *testing.T) {
		head := &list.Node[int]{Value: 1}

		want := &list.Node[int]{Value: 1}

		rev(head, 3)

		assertEqualLists(t, head, want)
	})

	t.Run("it should not mutate the list if k < 2", func(t *testing.T) {
		head := &list.Node[int]{Value: 1}
		head.Next = &list.Node[int]{Value: 2}
		head.Next.Next = &list.Node[int]{Value: 3}

		want := &list.Node[int]{Value: 1}
		want.Next = &list.Node[int]{Value: 2}
		want.Next.Next = &list.Node[int]{Value: 3}

		rev(head, 1)

		assertEqualLists(t, head, want)
	})

	t.Run("it should reverse alternating k sub-lists", func(t *testing.T) {
		head := &list.Node[int]{Value: 1}
		head.Next = &list.Node[int]{Value: 2}
		head.Next.Next = &list.Node[int]{Value: 3}
		head.Next.Next.Next = &list.Node[int]{Value: 4}
		head.Next.Next.Next.Next = &list.Node[int]{Value: 5}
		head.Next.Next.Next.Next.Next = &list.Node[int]{Value: 6}
		head.Next.Next.Next.Next.Next.Next = &list.Node[int]{Value: 7}
		head.Next.Next.Next.Next.Next.Next.Next = &list.Node[int]{Value: 8}
		head.Next.Next.Next.Next.Next.Next.Next.Next = &list.Node[int]{Value: 9}
		head.Next.Next.Next.Next.Next.Next.Next.Next.Next = &list.Node[int]{Value: 10}
		head.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next = &list.Node[int]{Value: 11}

		want := &list.Node[int]{Value: 3}
		want.Next = &list.Node[int]{Value: 2}
		want.Next.Next = &list.Node[int]{Value: 1}
		want.Next.Next.Next = &list.Node[int]{Value: 4}
		want.Next.Next.Next.Next = &list.Node[int]{Value: 5}
		want.Next.Next.Next.Next.Next = &list.Node[int]{Value: 6}
		want.Next.Next.Next.Next.Next.Next = &list.Node[int]{Value: 9}
		want.Next.Next.Next.Next.Next.Next.Next = &list.Node[int]{Value: 8}
		want.Next.Next.Next.Next.Next.Next.Next.Next = &list.Node[int]{Value: 7}
		want.Next.Next.Next.Next.Next.Next.Next.Next.Next = &list.Node[int]{Value: 10}
		want.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next = &list.Node[int]{Value: 11}

		got := rev(head, 3)

		assertEqualLists(t, got, want)
	})
}

func assertEqualLists(t *testing.T, got, want *list.Node[int]) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
