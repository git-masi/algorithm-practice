package main

import (
	"reflect"
	"testing"

	tree "github.com/git-masi/algorithm-practice/go/data-structures/binary-tree"
)

func TestTravZigZag(t *testing.T) {
	t.Run("It should return a 2D slice with 1 sub-slice and 1 value in the sub-slice", func(t *testing.T) {
		bt := &tree.Node{Value: 1}
		got := travZigZag(bt)
		want := [][]int{{1}}

		assertEqualSlices(t, got, want)
	})

	t.Run("It should return an empty slice if the input tree is nil", func(t *testing.T) {
		var bt *tree.Node
		got := travZigZag(bt)
		want := [][]int{}

		assertEqualSlices(t, got, want)
	})

	t.Run("It should return values in zig zag order", func(t *testing.T) {
		bt := &tree.Node{Value: 1}
		bt.Left = &tree.Node{Value: 2}
		bt.Right = &tree.Node{Value: 3}
		bt.Left.Left = &tree.Node{Value: 4}
		bt.Left.Right = &tree.Node{Value: 5}
		bt.Right.Left = &tree.Node{Value: 6}
		bt.Right.Right = &tree.Node{Value: 7}

		got := travZigZag(bt)
		want := [][]int{{1}, {3, 2}, {4, 5, 6, 7}}

		assertEqualSlices(t, got, want)
	})
}

func assertEqualSlices(t *testing.T, got, want [][]int) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
