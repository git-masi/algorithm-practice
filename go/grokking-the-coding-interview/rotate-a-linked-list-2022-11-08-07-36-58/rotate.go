package main

import (
	list "github.com/git-masi/algorithm-practice/go/data-structures/linked-list"
)

func rotate(head *list.Node[int], k int) *list.Node[int] {
	if head == nil || head.Next == nil || k < 2 {
		return head
	}

	var prev, node, tail, newHead *list.Node[int]

	node = head

	for i := 1; i <= k; i++ {
		if node == nil {
			tail = prev
			node = head
		}

		if i == k {
			prev.Next = nil
		}

		prev = node
		node = node.Next
	}

	if prev == head {
		return head
	} else {
		newHead = prev
	}

	if tail != nil {
		tail.Next = head
		return newHead
	}

	for {
		if node == nil {
			break
		}

		prev = node
		node = node.Next
	}

	prev.Next = head

	return newHead
}
