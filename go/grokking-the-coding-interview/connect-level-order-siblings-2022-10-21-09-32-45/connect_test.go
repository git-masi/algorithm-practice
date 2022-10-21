package main

import (
	"reflect"
	"testing"

	tree "github.com/git-masi/algorithm-practice/go/data-structures/binary-tree"
)

func TestConnect(t *testing.T) {
	t.Run("It should not mutate a tree with 1 node", func(t *testing.T) {
		got := &tree.Node{Value: 1}
		want := &tree.Node{Value: 1}
		connect(got)

		assertEqualTrees(t, got, want)
	})

	t.Run("There should be 2 root nodes", func(t *testing.T) {
		one := &tree.Node{Value: 1}
		two := &tree.Node{Value: 2}
		three := &tree.Node{Value: 3}

		one.Left = two
		one.Right = three

		expectOne := &tree.Node{Value: 1}
		expectTwo := &tree.Node{Value: 2}

		expectTwo.Right = three

		connect(one)

		assertEqualTrees(t, one, expectOne)
		assertEqualTrees(t, two, expectTwo)
	})

	t.Run("There should be 4 root nodes with no children", func(t *testing.T) {
		one := &tree.Node{Value: 1}
		two := &tree.Node{Value: 2}
		three := &tree.Node{Value: 3}
		four := &tree.Node{Value: 4}

		one.Right = two
		one.Right.Right = three
		one.Right.Right.Right = four

		expectOne := &tree.Node{Value: 1}
		expectTwo := &tree.Node{Value: 2}
		expectThree := &tree.Node{Value: 3}
		expectFour := &tree.Node{Value: 4}

		connect(one)

		assertEqualTrees(t, one, expectOne)
		assertEqualTrees(t, two, expectTwo)
		assertEqualTrees(t, three, expectThree)
		assertEqualTrees(t, four, expectFour)
	})

	t.Run("There should be 4 root nodes: 2 with children, 2 without", func(t *testing.T) {
		one := &tree.Node{Value: 2}
		two := &tree.Node{Value: 4}
		three := &tree.Node{Value: 6}
		four := &tree.Node{Value: 8}
		five := &tree.Node{Value: 10}
		six := &tree.Node{Value: 12}

		one.Right = two
		one.Right.Left = three
		one.Right.Right = four
		one.Right.Left.Left = five
		one.Right.Right.Left = six

		expectOne := &tree.Node{Value: 2}
		expectTwo := &tree.Node{Value: 4}
		expectThree := &tree.Node{Value: 6}
		expectThree.Right = four
		expectFive := &tree.Node{Value: 10}
		expectFive.Right = six

		connect(one)

		assertEqualTrees(t, one, expectOne)
		assertEqualTrees(t, two, expectTwo)
		assertEqualTrees(t, three, expectThree)
		assertEqualTrees(t, five, expectFive)
	})
}

func assertEqualTrees(t *testing.T, got, want *tree.Node) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
