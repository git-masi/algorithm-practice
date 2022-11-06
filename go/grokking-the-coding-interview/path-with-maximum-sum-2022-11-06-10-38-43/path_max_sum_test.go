package main

import (
	"reflect"
	"testing"

	tree "github.com/git-masi/algorithm-practice/go/data-structures/binary-tree"
)

func TestPathMaxSum(t *testing.T) {
	t.Run("it should return an empty slice given a nil root", func(t *testing.T) {
		got := pathMaxSum(nil)
		want := []*tree.Node{}

		assertEqualSlices(t, got, want)
	})

	t.Run("it should return a slice with 1 node given a tree with 1 node", func(t *testing.T) {
		root := &tree.Node{Value: 1}
		got := pathMaxSum(root)
		want := []*tree.Node{root}

		assertEqualSlices(t, got, want)
	})

	t.Run("it should return a slice with multiple nodes including the root", func(t *testing.T) {
		root := &tree.Node{Value: 1}
		root.Left = &tree.Node{Value: 2}
		root.Left.Left = &tree.Node{Value: 3}
		root.Left.Right = &tree.Node{Value: 4}
		root.Right = &tree.Node{Value: 5}
		root.Right.Left = &tree.Node{Value: 7}
		root.Right.Right = &tree.Node{Value: 6}
		got := pathMaxSum(root)
		want := []*tree.Node{root, root.Left, root.Left.Right, root.Right, root.Right.Left}

		assertEqualSlices(t, got, want)
	})

	t.Run("it should return a slice with 3 nodes, the root of the path is not the root node of the tree", func(t *testing.T) {
		root := &tree.Node{Value: 1}
		root.Left = &tree.Node{Value: 2}
		root.Right = &tree.Node{Value: 3}
		root.Right.Left = &tree.Node{Value: 22}
		root.Right.Right = &tree.Node{Value: 4}
		got := pathMaxSum(root)
		want := []*tree.Node{root.Right, root.Right.Left, root.Right.Right}

		assertEqualSlices(t, got, want)
	})
}

func assertEqualSlices(t *testing.T, got, want []*tree.Node) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
