package main

import "fmt"

// 206. Reverse Linked List

type listNode struct {
	Val  int
	Next *listNode
}

func main() {
	node1 := &listNode{
		Val: 1,
	}
	node2 := &listNode{
		Val: 2,
	}
	node3 := &listNode{
		Val: 3,
	}
	node4 := &listNode{
		Val: 4,
	}
	node5 := &listNode{
		Val: 5,
	}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	revHead := reverseList(node1)
	fmt.Println(revHead)
}

/**
 * Definition for singly-linked list.
 * type listNode struct {
 *     Val int
 *     Next *listNode
 * }
 */
func reverseList(head *listNode) *listNode {
	if head == nil {
		return head
	}
	if head.Next == nil {
		return head
	}

	var revHead *listNode
	for head != nil {
		tmp := head.Next
		head.Next = revHead
		revHead = head
		head = tmp
	}
	return revHead
}
