package main

import (
	"testing"

	tree "github.com/git-masi/algorithm-practice/go/data-structures/binary-tree"
)

func TestDiameter(t *testing.T) {
	t.Run("it should return 3 for the smallest possible tree", func(t *testing.T) {
		root := &tree.Node{Value: 1}
		root.Left = &tree.Node{Value: 2}
		root.Right = &tree.Node{Value: 3}
		got := diameter(root)
		want := 3

		assertEqualInt(t, got, want)
	})

	t.Run("it should return 5", func(t *testing.T) {
		// complete balanced tree
		root := &tree.Node{Value: 1}
		root.Left = &tree.Node{Value: 2}
		root.Left.Left = &tree.Node{Value: 3}
		root.Left.Right = &tree.Node{Value: 4}
		root.Right = &tree.Node{Value: 5}
		root.Right.Left = &tree.Node{Value: 6}
		root.Right.Right = &tree.Node{Value: 7}
		got := diameter(root)
		want := 5

		assertEqualInt(t, got, want)
	})

	t.Run("it should return 7, diameter does not pass through root", func(t *testing.T) {
		root := &tree.Node{Value: 1}
		root.Left = &tree.Node{Value: 2}
		// 3 is the root of the diameter
		root.Right = &tree.Node{Value: 3}
		root.Right.Left = &tree.Node{Value: 4}
		root.Right.Left.Left = &tree.Node{Value: 5}
		root.Right.Left.Left.Left = &tree.Node{Value: 6}
		root.Right.Right = &tree.Node{Value: 7}
		root.Right.Right.Right = &tree.Node{Value: 8}
		root.Right.Right.Right.Right = &tree.Node{Value: 9}
		got := diameter(root)
		want := 7

		assertEqualInt(t, got, want)
	})
}

func assertEqualInt(t *testing.T, got, want int) {
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
