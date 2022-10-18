package main

import (
	"reflect"
	"testing"

	binarytree "github.com/git-masi/algorithm-practice/go/data-structures/binary-tree"
)

func TestTravLevelOrder(t *testing.T) {
	t.Run("It should return the correct values for a complete tree", func(t *testing.T) {
		bt := binarytree.Node{Value: 1}
		bt.Left = &binarytree.Node{Value: 2}
		bt.Right = &binarytree.Node{Value: 3}
		bt.Left.Left = &binarytree.Node{Value: 4}
		bt.Left.Right = &binarytree.Node{Value: 5}
		bt.Right.Left = &binarytree.Node{Value: 6}
		bt.Right.Right = &binarytree.Node{Value: 7}

		got := travLevelOrder(bt)

		want := LevelOrderValues{[]int{1}, []int{2, 3}, []int{4, 5, 6, 7}}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("It should return the correct values for a lopsided tree", func(t *testing.T) {
		bt := binarytree.Node{Value: 1}
		bt.Right = &binarytree.Node{Value: 2}
		bt.Right.Right = &binarytree.Node{Value: 3}
		bt.Right.Right.Right = &binarytree.Node{Value: 4}

		got := travLevelOrder(bt)

		want := LevelOrderValues{[]int{1}, []int{2}, []int{3}, []int{4}}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
