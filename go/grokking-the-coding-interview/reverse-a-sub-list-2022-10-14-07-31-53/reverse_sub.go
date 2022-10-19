package main

import linkedlist "github.com/git-masi/algorithm-practice/go/data-structures/linked-list"

func reverseSub(ll *linkedlist.LinkedList[int], p, q int) {
	if ll.Size < 2 || p == q {
		return
	}

	var prev *linkedlist.Node[int]
	var beforeSubList *linkedlist.Node[int]
	var subListStart *linkedlist.Node[int]
	curr := ll.Head

	for {
		if curr.Value == p || curr == nil {
			break
		}

		prev = curr
		curr = curr.Next
	}

	beforeSubList = prev
	subListStart = curr

	for {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next

		if prev.Value == q || curr == nil {
			break
		}
	}

	if beforeSubList != nil {
		// prev is the last node we processed
		// This also means the head hasn't changed
		beforeSubList.Next = prev
	} else {
		// If there was no node before the start of the sub list that means the
		// first node in the sub list was the head. In which case we need to update
		// the head of the linked list to be the last node in the sub list that we
		// processed
		ll.Head = prev
	}

	// The current value is the node after the last node we processed
	// The old sub list start should be the last node in the sub list
	subListStart.Next = curr

	if curr == nil {
		// If the current node is nil we've reached the end of the list
		// which means that the node we started with is the new tail
		ll.Tail = subListStart
	}
}
