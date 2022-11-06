package main

import (
	"testing"

	tree "github.com/git-masi/algorithm-practice/go/data-structures/binary-tree"
)

func TestCountPaths(t *testing.T) {
	t.Run("it should return 0 given a nil root", func(t *testing.T) {
		got := countPaths(nil, 0)
		want := 0

		assertEqualInt(t, got, want)
	})

	t.Run("it should return 1 given a single node equal to target", func(t *testing.T) {
		root := &tree.Node{Value: 1}
		got := countPaths(root, 1)
		want := 1

		assertEqualInt(t, got, want)
	})

	t.Run("it should return 0 given a single node not equal to target", func(t *testing.T) {
		root := &tree.Node{Value: 1}
		got := countPaths(root, 2)
		want := 0

		assertEqualInt(t, got, want)
	})

	t.Run("it should return 1 given multiple nodes and 1 branch equal to target", func(t *testing.T) {
		root := &tree.Node{Value: 3}
		root.Left = &tree.Node{Value: 2}
		root.Right = &tree.Node{Value: 1}
		got := countPaths(root, 5)
		want := 1

		assertEqualInt(t, got, want)
	})

	t.Run("it should return 0 given multiple nodes where no branch equals the target", func(t *testing.T) {
		root := &tree.Node{Value: 3}
		root.Left = &tree.Node{Value: 2}
		root.Right = &tree.Node{Value: 1}
		got := countPaths(root, 99)
		want := 0

		assertEqualInt(t, got, want)
	})

	t.Run("it should return 4", func(t *testing.T) {
		root := &tree.Node{Value: 1}
		root.Left = &tree.Node{Value: 2}
		root.Left.Left = &tree.Node{Value: 2}
		root.Right = &tree.Node{Value: 4}
		root.Right.Left = &tree.Node{Value: 1}
		root.Right.Right = &tree.Node{Value: 1}
		got := countPaths(root, 5)
		want := 4

		assertEqualInt(t, got, want)
	})

	t.Run("it should return 3", func(t *testing.T) {
		root := &tree.Node{Value: 5}
		root.Left = &tree.Node{Value: 1}
		root.Right = &tree.Node{Value: 2}
		root.Right.Right = &tree.Node{Value: 2}
		root.Right.Right.Right = &tree.Node{Value: 2}
		root.Right.Right.Right.Right = &tree.Node{Value: 3}
		root.Right.Right.Right.Right.Right = &tree.Node{Value: 5}
		root.Right.Right.Right.Right.Right.Right = &tree.Node{Value: 6}
		got := countPaths(root, 5)
		want := 3

		assertEqualInt(t, got, want)
	})
}

func assertEqualInt(t *testing.T, got, want int) {
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
