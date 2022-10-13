package main

import linkedlist "github.com/git-masi/algorithm-practice/go/data-structures/linked-list"

func reverse(ll *linkedlist.LinkedList) {
	if ll.Head == nil {
		return
	}

	var prev *linkedlist.LinkedListNode
	curr := ll.Head

	ll.Tail = curr

	for {
		if curr == nil {
			break
		}

		temp := curr.Next

		curr.Next = prev
		prev = curr
		curr = temp
	}

	ll.Head = prev
}
