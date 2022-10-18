package main

import (
	"reflect"
	"testing"

	binarytree "github.com/git-masi/algorithm-practice/go/data-structures/binary-tree"
)

func TestTravLevelOrder(t *testing.T) {
	t.Run("It should return a single sub slice for a binary tree consisting of one node", func(t *testing.T) {
		bt := binarytree.Node{Value: 1}
		got := travLevelOrder(bt)
		want := LevelOrderValues{[]int{1}}

		assertEqualValues(t, got, want)
	})

	t.Run("It should return two sub slices for a binary tree consisting of three nodes", func(t *testing.T) {
		bt := binarytree.Node{Value: 1}
		bt.Left = &binarytree.Node{Value: 2}
		bt.Right = &binarytree.Node{Value: 3}

		got := travLevelOrder(bt)
		want := LevelOrderValues{[]int{1}, []int{2, 3}}

		assertEqualValues(t, got, want)
	})

	t.Run("It should return the correct values for an incomplete tree", func(t *testing.T) {
		// Level 1
		bt := binarytree.Node{Value: 1}
		// Level 2
		bt.Left = &binarytree.Node{Value: 2}
		bt.Right = &binarytree.Node{Value: 3}
		// Level 3
		bt.Left.Right = &binarytree.Node{Value: 4}
		bt.Right.Left = &binarytree.Node{Value: 5}
		bt.Right.Right = &binarytree.Node{Value: 6}

		got := travLevelOrder(bt)
		want := LevelOrderValues{[]int{1}, []int{2, 3}, []int{4, 5, 6}}

		assertEqualValues(t, got, want)
	})

	t.Run("It should return the correct values for a lopsided tree", func(t *testing.T) {
		// Level 1
		bt := binarytree.Node{Value: 1}
		// Level 2
		bt.Right = &binarytree.Node{Value: 2}
		// Level 3
		bt.Right.Right = &binarytree.Node{Value: 3}
		// Level 4
		bt.Right.Right.Right = &binarytree.Node{Value: 4}
		// Level 5
		bt.Right.Right.Right.Right = &binarytree.Node{Value: 5}
		// Level 6
		bt.Right.Right.Right.Right.Right = &binarytree.Node{Value: 6}

		got := travLevelOrder(bt)
		want := LevelOrderValues{[]int{1}, []int{2}, []int{3}, []int{4}, []int{5}, []int{6}}

		assertEqualValues(t, got, want)
	})
}

func assertEqualValues(t *testing.T, got, want LevelOrderValues) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
