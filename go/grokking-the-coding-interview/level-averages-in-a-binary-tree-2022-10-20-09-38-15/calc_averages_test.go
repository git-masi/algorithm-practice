package main

import (
	"reflect"
	"testing"

	tree "github.com/git-masi/algorithm-practice/go/data-structures/binary-tree"
)

func TestCalcAverages(t *testing.T) {
	t.Run("It should return an empty slice if given a nil binary tree", func(t *testing.T) {
		got := calcAverages(nil)

		want := []float64{}

		assertEqualSlices(t, got, want)
	})

	t.Run("It should return a slice with 1 value", func(t *testing.T) {
		bt := &tree.Node{Value: 1}

		got := calcAverages(bt)

		want := []float64{1}

		assertEqualSlices(t, got, want)
	})

	t.Run("It should return a slice with 3 values", func(t *testing.T) {
		bt := &tree.Node{Value: 1}
		bt.Left = &tree.Node{Value: 2}
		bt.Right = &tree.Node{Value: 3}
		bt.Left.Left = &tree.Node{Value: 4}
		bt.Left.Right = &tree.Node{Value: 4}
		bt.Right.Left = &tree.Node{Value: 4}
		bt.Right.Right = &tree.Node{Value: 4}

		got := calcAverages(bt)

		want := []float64{1, 2.5, 4}

		assertEqualSlices(t, got, want)
	})
}

func assertEqualSlices(t *testing.T, got, want []float64) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
