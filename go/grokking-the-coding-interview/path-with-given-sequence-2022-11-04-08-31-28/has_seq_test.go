package main

import (
	"testing"

	tree "github.com/git-masi/algorithm-practice/go/data-structures/binary-tree"
)

func TestHasSeq(t *testing.T) {
	t.Run("it should return false for a 'nil' root", func(t *testing.T) {
		got := hasSeq(nil, nil)
		want := false

		assertEqualBool(t, got, want)
	})

	t.Run("it should return false for an empty sequence", func(t *testing.T) {
		root := &tree.Node{Value: 1}
		seq := []int{}
		got := hasSeq(root, seq)
		want := false

		assertEqualBool(t, got, want)
	})

	t.Run("it should return false for a nil sequence", func(t *testing.T) {
		root := &tree.Node{Value: 1}
		got := hasSeq(root, nil)
		want := false

		assertEqualBool(t, got, want)
	})

	t.Run("it should return true if 1 node matches a sequence of 1 element", func(t *testing.T) {
		root := &tree.Node{Value: 1}
		seq := []int{1}
		got := hasSeq(root, seq)
		want := true

		assertEqualBool(t, got, want)
	})

	t.Run("it should return false if 1 node doesn't match a sequence of 1 element", func(t *testing.T) {
		root := &tree.Node{Value: 1}
		seq := []int{2}
		got := hasSeq(root, seq)
		want := false

		assertEqualBool(t, got, want)
	})

	t.Run("it should return true: multiple nodes, multiple items in sequence", func(t *testing.T) {
		root := &tree.Node{Value: 1}
		root.Left = &tree.Node{Value: 2}
		root.Right = &tree.Node{Value: 3}
		root.Right.Left = &tree.Node{Value: 4}
		seq := []int{1, 3, 4}
		got := hasSeq(root, seq)
		want := true

		assertEqualBool(t, got, want)
	})

	t.Run("it should return false: multiple nodes, multiple items in sequence", func(t *testing.T) {
		root := &tree.Node{Value: 1}
		root.Left = &tree.Node{Value: 2}
		root.Right = &tree.Node{Value: 2}
		root.Left.Left = &tree.Node{Value: 7}
		root.Left.Right = &tree.Node{Value: 8}
		root.Right.Left = &tree.Node{Value: 9}
		seq := []int{1, 2}
		got := hasSeq(root, seq)
		want := false

		assertEqualBool(t, got, want)
	})
}

func assertEqualBool(t *testing.T, got, want bool) {
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
