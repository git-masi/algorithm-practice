package main

import (
	"reflect"
	"testing"

	tree "github.com/git-masi/algorithm-practice/go/data-structures/binary-tree"
)

func TestRightView(t *testing.T) {
	t.Run("it should return an empty slice if given a 'nil' tree", func(t *testing.T) {
		got := rightView(nil)
		want := []*tree.Node{}

		assertEqualSlices(t, got, want)
	})

	t.Run("it should return a slice with 1 node", func(t *testing.T) {
		root := &tree.Node{Value: 1}
		got := rightView(root)
		want := []*tree.Node{root}

		assertEqualSlices(t, got, want)
	})

	t.Run("it should return a slice with 2 nodes", func(t *testing.T) {
		root := &tree.Node{Value: 1}
		root.Left = &tree.Node{Value: 2}
		root.Right = &tree.Node{Value: 3}

		got := rightView(root)
		want := []*tree.Node{root, root.Right}

		assertEqualSlices(t, got, want)
	})

	t.Run("it should return a slice with 4 nodes", func(t *testing.T) {
		root := &tree.Node{Value: 1}
		root.Left = &tree.Node{Value: 2}
		root.Left.Right = &tree.Node{Value: 3}
		root.Left.Right.Left = &tree.Node{Value: 4}
		root.Left.Right.Right = &tree.Node{Value: 5}

		got := rightView(root)
		want := []*tree.Node{root, root.Left, root.Left.Right, root.Left.Right.Right}

		assertEqualSlices(t, got, want)
	})
}

func assertEqualSlices(t testing.TB, got, want []*tree.Node) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
