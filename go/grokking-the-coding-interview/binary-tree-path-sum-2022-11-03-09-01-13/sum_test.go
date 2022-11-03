package main

import (
	"testing"

	tree "github.com/git-masi/algorithm-practice/go/data-structures/binary-tree"
)

func TestSum(t *testing.T) {
	t.Run("it should return false if the btree is 'nil'", func(t *testing.T) {
		got := sum(nil, 0)
		want := false

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("it should return 'true' if there is 1 node and the target equals the node's value", func(t *testing.T) {
		got := sum(&tree.Node{Value: 1}, 1)
		want := true

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("it should return 'false' if there are multiple nodes that don't match the target", func(t *testing.T) {
		root := &tree.Node{Value: 1}
		root.Left = &tree.Node{Value: 2}
		root.Right = &tree.Node{Value: 3}
		got := sum(root, 27)
		want := false

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("it should return 'true' if there are multiple nodes that match the target", func(t *testing.T) {
		root := &tree.Node{Value: 1}
		root.Left = &tree.Node{Value: 2}
		root.Left.Left = &tree.Node{Value: 3}
		root.Left.Right = &tree.Node{Value: 4}
		root.Left.Left.Left = &tree.Node{Value: 5}
		got := sum(root, 7)
		want := true

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
