package main

import (
	"reflect"
	"testing"

	tree "github.com/git-masi/algorithm-practice/go/data-structures/binary-tree"
)

func TestConnectSiblings(t *testing.T) {
	t.Run("It should not mutate a tree with a single node", func(t *testing.T) {
		bt := &tree.Node{Value: 1}
		want := &tree.Node{Value: 1}

		connectSiblings(bt)

		assertEqualTrees(t, bt, want)
	})

	t.Run("It should connect sibling nodes", func(t *testing.T) {
		bt := &tree.Node{Value: 1}
		bt.Left = &tree.Node{Value: 2}
		bt.Right = &tree.Node{Value: 3}
		bt.Left.Left = &tree.Node{Value: 4}
		bt.Left.Right = &tree.Node{Value: 5}
		bt.Right.Left = &tree.Node{Value: 6}
		bt.Right.Right = &tree.Node{Value: 7}

		want := &tree.Node{Value: 1}
		want.Right = &tree.Node{Value: 2}
		want.Right.Right = &tree.Node{Value: 3}
		want.Right.Right.Right = &tree.Node{Value: 4}
		want.Right.Right.Right.Right = &tree.Node{Value: 5}
		want.Right.Right.Right.Right.Right = &tree.Node{Value: 6}
		want.Right.Right.Right.Right.Right.Right = &tree.Node{Value: 7}

		connectSiblings(bt)

		assertEqualTrees(t, bt, want)
	})
}

func assertEqualTrees(t *testing.T, got, want *tree.Node) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
