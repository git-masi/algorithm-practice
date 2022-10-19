package main

import (
	"reflect"
	"testing"

	tree "github.com/git-masi/algorithm-practice/go/data-structures/binary-tree"
	list "github.com/git-masi/algorithm-practice/go/data-structures/linked-list"
)

func TestRevLevelOrder(t *testing.T) {
	t.Run("It should return a linked list with one node", func(t *testing.T) {
		bt := &tree.Node{Value: 0}
		got := revLevelOrder(bt)
		want := list.LinkedList[[]int]{}

		want.Push([]int{0})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("It should return a linked list with three nodes", func(t *testing.T) {
		bt := &tree.Node{Value: 0}
		bt.Left = &tree.Node{Value: 1}
		bt.Right = &tree.Node{Value: 2}
		bt.Left.Left = &tree.Node{Value: 3}
		bt.Left.Right = &tree.Node{Value: 4}
		bt.Right.Left = &tree.Node{Value: 5}
		bt.Right.Right = &tree.Node{Value: 6}

		got := revLevelOrder(bt)
		want := list.LinkedList[[]int]{}

		want.Push([]int{3, 4, 5, 6})
		want.Push([]int{1, 2})
		want.Push([]int{0})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
