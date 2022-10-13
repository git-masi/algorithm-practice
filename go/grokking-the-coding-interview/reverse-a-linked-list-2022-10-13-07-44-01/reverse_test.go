package main

import (
	"reflect"
	"testing"

	linkedlist "github.com/git-masi/algorithm-practice/go/data-structures/linked-list"
)

func TestReverse(t *testing.T) {
	t.Run("It should not mutate and empty linked list", func(t *testing.T) {
		ll := createTestList(0)

		reverse(&ll)

		want := createTestList(0)

		assertListsAreEqual(t, ll, want)
	})

	t.Run("It should reverse the head and the tail", func(t *testing.T) {
		ll := createTestList(2)

		reverse(&ll)

		want := createTestList(-2)

		assertListsAreEqual(t, ll, want)
	})

	t.Run("It should reverse a long linked list", func(t *testing.T) {
		ll := createTestList(6)

		reverse(&ll)

		want := createTestList(-6)

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
		t.Errorf("got head %v tail %v, want head %v tail %v", got.Head, got.Tail, want.Head, want.Tail)
	}
}
