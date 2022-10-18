package binarytree

import "testing"

func TestNode(t *testing.T) {
	t.Run("It should have zero values by default", func(t *testing.T) {
		node := Node{}

		if node.Left != nil {
			t.Errorf("node.Left should be 'nil', got %v", node.Left)
		}

		if node.Right != nil {
			t.Errorf("node.Right should be 'nil', got %v", node.Right)
		}

		if node.Value != 0 {
			t.Errorf("node.Value should be '0', got %d", node.Value)
		}
	})
}
