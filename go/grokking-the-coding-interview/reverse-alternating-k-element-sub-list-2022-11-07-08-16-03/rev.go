package main

import (
	list "github.com/git-masi/algorithm-practice/go/data-structures/linked-list"
)

func rev(head *list.Node[int], k int) *list.Node[int] {
	if k < 2 || head.Next == nil {
		return head
	}

	node := head
	var prev, next, before, first, last *list.Node[int]

	for {
		if node == nil {
			break
		}

		i, j := 0, 0
		first = node

		for {
			if i == k || node == nil {
				break
			}

			next = node.Next
			node.Next = prev
			prev = node
			node = next

			i++
		}

		last = prev

		if before != nil {
			before.Next = last
		} else {
			head = last
		}

		first.Next = node

		for {
			if j == k || node == nil {
				break
			}

			next = node.Next
			prev = node
			node = next

			j++
		}

		before = prev
	}

	return head
}
