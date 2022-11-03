package main

import (
	"reflect"
	"testing"

	tree "github.com/git-masi/algorithm-practice/go/data-structures/binary-tree"
)

func TestSum(t *testing.T) {
	t.Run("it should return an empty slice given a nil tree", func(t *testing.T) {
		got := sum(nil, 0)
		want := [][]*tree.Node{}

		assertEqualSlices(t, got, want)
	})

	t.Run("it should return a slice with 1 path", func(t *testing.T) {
		root := &tree.Node{Value: 1}
		got := sum(root, 1)
		want := [][]*tree.Node{{root}}

		assertEqualSlices(t, got, want)
	})

	t.Run("it should return a slice with 0 paths", func(t *testing.T) {
		got := sum(&tree.Node{Value: 1}, 2)
		want := [][]*tree.Node{}

		assertEqualSlices(t, got, want)
	})

	t.Run("it should return a slice with 1 path and 2 nodes", func(t *testing.T) {
		root := &tree.Node{Value: 1}
		root.Left = &tree.Node{Value: 2}
		root.Right = &tree.Node{Value: 3}
		got := sum(root, 3)
		want := [][]*tree.Node{{root, root.Left}}

		assertEqualSlices(t, got, want)
	})

	t.Run("it should return a slice with 1 path and 3 nodes", func(t *testing.T) {
		root := &tree.Node{Value: 1}
		root.Left = &tree.Node{Value: 2}
		root.Right = &tree.Node{Value: 5}
		root.Left.Left = &tree.Node{Value: 3}
		root.Left.Right = &tree.Node{Value: 4}
		got := sum(root, 7)
		want := [][]*tree.Node{{root, root.Left, root.Left.Right}}

		assertEqualSlices(t, got, want)
	})

	t.Run("it should return 2 slices", func(t *testing.T) {
		root := &tree.Node{Value: 1}
		root.Left = &tree.Node{Value: 2}
		root.Right = &tree.Node{Value: 5}
		root.Left.Right = &tree.Node{Value: 9}
		root.Right.Left = &tree.Node{Value: 6}
		got := sum(root, 12)
		want := [][]*tree.Node{{root, root.Left, root.Left.Right}, {root, root.Right, root.Right.Left}}

		assertEqualSlices(t, got, want)
	})
}

func assertEqualSlices(t *testing.T, got, want [][]*tree.Node) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
