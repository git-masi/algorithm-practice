package main

import (
	"testing"

	tree "github.com/git-masi/algorithm-practice/go/data-structures/binary-tree"
)

func TestSuccessor(t *testing.T) {
	t.Run("It should return 'nil' for a binary tree with 1 node", func(t *testing.T) {
		bt := &tree.Node{Value: 1}
		target := bt
		got := successor(bt, target)
		var want *tree.Node = nil

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("It should return 2", func(t *testing.T) {
		bt := &tree.Node{Value: 1}
		target := &tree.Node{Value: 2}
		bt.Left = target
		want := &tree.Node{Value: 3}
		bt.Right = want
		got := successor(bt, target)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("It should return 6", func(t *testing.T) {
		bt := &tree.Node{Value: 1}
		bt.Left = &tree.Node{Value: 2}
		bt.Right = &tree.Node{Value: 3}
		bt.Left.Left = &tree.Node{Value: 4}
		target := &tree.Node{Value: 5}
		bt.Left.Right = target
		want := &tree.Node{Value: 6}
		bt.Right.Left = want
		bt.Right.Right = &tree.Node{Value: 7}
		got := successor(bt, target)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
