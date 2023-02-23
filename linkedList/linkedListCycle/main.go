package main

import "fmt"

// 141. Linked List Cycle

type listNode struct {
	Val  int
	Next *listNode
}

func main() {
	node1 := &listNode{
		Val: 3,
	}
	node2 := &listNode{
		Val: 2,
	}
	node3 := &listNode{
		Val: 0,
	}
	node4 := &listNode{
		Val: 4,
	}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2

	fmt.Println(hasCycle(node1))
}


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *listNode) bool {
	if head == nil {
		return false
	}
	if head.Next == nil {
		return false
	}

	slowPointer := head
	fastPointer := head
	for ; fastPointer.Next != nil; {
		slowPointer = slowPointer.Next
		fastPointer = fastPointer.Next.Next
		if fastPointer == nil {
			return false
		}
		if slowPointer == fastPointer {
			return true
		}
	}
	return false
}