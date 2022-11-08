package main

import (
	"reflect"
	"testing"

	list "github.com/git-masi/algorithm-practice/go/data-structures/linked-list"
)

func TestRotate(t *testing.T) {
	t.Run("it should not mutate the list if there is only 1 node", func(t *testing.T) {
		head := &list.Node[int]{Value: 1}
		want := head
		k := 2

		got := rotate(head, k)

		assertEqualLists(t, got, want)
	})

	t.Run("it should not mutate the list if k < 2", func(t *testing.T) {
		head := &list.Node[int]{Value: 1}
		want := head
		k := 1

		got := rotate(head, k)

		assertEqualLists(t, got, want)
	})

	t.Run("it should return 2 as the head node if k = 2 and list len = 2", func(t *testing.T) {
		head := &list.Node[int]{Value: 1}
		head.Next = &list.Node[int]{Value: 2}

		want := &list.Node[int]{Value: 2}
		want.Next = &list.Node[int]{Value: 1}

		k := 2

		got := rotate(head, k)

		assertEqualLists(t, got, want)
	})

	t.Run("it should return 5 as the head node if k = 10 and list len = 5", func(t *testing.T) {
		head := &list.Node[int]{Value: 1}
		head.Next = &list.Node[int]{Value: 2}
		head.Next.Next = &list.Node[int]{Value: 3}
		head.Next.Next.Next = &list.Node[int]{Value: 4}
		head.Next.Next.Next.Next = &list.Node[int]{Value: 5}

		want := &list.Node[int]{Value: 5}
		want.Next = &list.Node[int]{Value: 1}
		want.Next.Next = &list.Node[int]{Value: 2}
		want.Next.Next.Next = &list.Node[int]{Value: 3}
		want.Next.Next.Next.Next = &list.Node[int]{Value: 4}

		k := 5

		got := rotate(head, k)

		assertEqualLists(t, got, want)
	})
}

func assertEqualLists(t *testing.T, got, want *list.Node[int]) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
