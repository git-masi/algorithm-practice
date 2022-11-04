package main

import (
	"testing"

	tree "github.com/git-masi/algorithm-practice/go/data-structures/binary-tree"
)

func TestSumPaths(t *testing.T) {
	t.Run("it should return 0 for a 'nil' btree", func(t *testing.T) {
		got := sumPaths(nil)
		want := 0

		assertEqualInts(t, got, want)
	})

	t.Run("it should return 1", func(t *testing.T) {
		root := &tree.Node{Value: 1}
		got := sumPaths(root)
		want := 1

		assertEqualInts(t, got, want)
	})

	t.Run("it should return 25", func(t *testing.T) {
		root := &tree.Node{Value: 1}
		root.Left = &tree.Node{Value: 2}
		root.Right = &tree.Node{Value: 3}
		got := sumPaths(root)
		want := 25

		assertEqualInts(t, got, want)
	})

	t.Run("it should return 24", func(t *testing.T) {
		root := &tree.Node{Value: 0}
		root.Left = &tree.Node{Value: 0}
		root.Right = &tree.Node{Value: 7}
		root.Left.Left = &tree.Node{Value: 9}
		root.Left.Right = &tree.Node{Value: 8}
		got := sumPaths(root)
		want := 24

		assertEqualInts(t, got, want)
	})

	t.Run("it should return 332", func(t *testing.T) {
		root := &tree.Node{Value: 1}
		root.Left = &tree.Node{Value: 0}
		root.Right = &tree.Node{Value: 1}
		root.Left.Left = &tree.Node{Value: 1}
		root.Right.Left = &tree.Node{Value: 6}
		root.Right.Right = &tree.Node{Value: 5}
		got := sumPaths(root)
		want := 332

		assertEqualInts(t, got, want)
	})
}

func assertEqualInts(t *testing.T, got, want int) {
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
