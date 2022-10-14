package main

import (
	"reflect"
	"testing"

	linkedlist "github.com/git-masi/algorithm-practice/go/data-structures/linked-list"
)

func TestReverseSub(t *testing.T) {
	t.Run("It should not mutate an empty list", func(t *testing.T) {
		ll := linkedlist.LinkedList{}

		reverseSub(&ll, 0, 1)

		want := linkedlist.LinkedList{}

		assertListsAreEqual(t, ll, want)
	})

	t.Run("It should not mutate a list with size < 2", func(t *testing.T) {
		ll := createTestList(1)

		reverseSub(&ll, 0, 1)

		want := createTestList(-1)

		assertListsAreEqual(t, ll, want)
	})

	t.Run("It should reverse the head and the tail", func(t *testing.T) {
		ll := createTestList(2)

		reverseSub(&ll, 0, 1)

		want := createTestList(-2)

		assertListsAreEqual(t, ll, want)
	})

	t.Run("It should reverse the second and third nodes", func(t *testing.T) {
		ll := createTestList(4)

		reverseSub(&ll, 1, 2)

		want := linkedlist.LinkedList{}
		want.Insert(0)
		want.Insert(2)
		want.Insert(1)
		want.Insert(3)

		assertListsAreEqual(t, ll, want)
	})
}

func createTestList(size int) linkedlist.LinkedList {
	ll := linkedlist.LinkedList{}

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

func assertListsAreEqual(t *testing.T, got, want linkedlist.LinkedList) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got values: %v, want values: %v", got.ToSlice(), want.ToSlice())
	}
}
