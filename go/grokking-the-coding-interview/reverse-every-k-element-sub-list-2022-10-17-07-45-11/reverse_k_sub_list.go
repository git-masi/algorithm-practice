package main

import linkedlist "github.com/git-masi/algorithm-practice/go/data-structures/linked-list"

func reverseSubList(head *linkedlist.Node[int], k int) *linkedlist.Node[int] {
	if k < 2 || head == nil {
		return head
	}

	var prev *linkedlist.Node[int]
	curr := head

	for {
		var next *linkedlist.Node[int]
		lastNodeOfPrevSubList := prev
		lastNodeOfSubList := curr
		count := 0

		for {
			if curr == nil || count >= k {
				break
			}

			next = curr.Next
			curr.Next = prev
			prev = curr
			curr = next
			count++
		}

		if lastNodeOfPrevSubList != nil {
			lastNodeOfPrevSubList.Next = prev
		} else {
			head = prev
		}

		lastNodeOfSubList.Next = curr

		if curr == nil {
			break
		}

		prev = lastNodeOfSubList
	}

	return head
}
