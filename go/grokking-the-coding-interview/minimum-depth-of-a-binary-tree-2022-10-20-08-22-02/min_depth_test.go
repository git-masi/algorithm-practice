package main

import (
	"testing"

	tree "github.com/git-masi/algorithm-practice/go/data-structures/binary-tree"
)

func TestMinDepth(t *testing.T) {
	t.Run("It should return 0 for a nil binary tree", func(t *testing.T) {
		got := minDepth(nil)
		want := 0

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("It should return 1 for a binary tree with a single node", func(t *testing.T) {
		bt := &tree.Node{Value: 1}
		got := minDepth(bt)
		want := 1

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("It should return 2 if there are 3 levels but there is a leaf node on level 2", func(t *testing.T) {
		bt := &tree.Node{Value: 1}
		bt.Left = &tree.Node{Value: 2}
		bt.Right = &tree.Node{Value: 3}
		bt.Right.Left = &tree.Node{Value: 4}
		bt.Right.Right = &tree.Node{Value: 5}

		got := minDepth(bt)
		want := 2

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("It should return 4 for an unbalanced tree with 4 levels", func(t *testing.T) {
		bt := &tree.Node{Value: 2}
		bt.Right = &tree.Node{Value: 4}
		bt.Right.Right = &tree.Node{Value: 6}
		bt.Right.Right.Right = &tree.Node{Value: 8}

		got := minDepth(bt)
		want := 4

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
