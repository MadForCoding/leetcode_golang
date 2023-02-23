package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 21. merge two sort list

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
		Val: 4,
	}
	node1.Next = node2
	node2.Next = node3

	node4 := &listNode{
		Val: 1,
	}
	node5 := &listNode{
		Val: 3,
	}
	node6 := &listNode{
		Val: 4,
	}
	node4.Next = node5
	node5.Next = node6

	head := mergeTwoListsSolution1(node1, node4)
	fmt.Println(printResult(head))
}


func printResult(head *listNode) string {
	if head == nil {
		return ""
	}
	builder := &strings.Builder{}
	for head != nil {
		builder.WriteString(strconv.FormatInt(int64(head.Val), 10))
		head = head.Next
		if head != nil {
			builder.WriteString(",")
		}
	}
	return builder.String()
}

// mergeTwoListsSolution1 - 只在原来的list上操作
func mergeTwoListsSolution1(list1 *listNode, list2 *listNode) *listNode {
	if list2 == nil {
		return list1
	}
	if list1 == nil {
		return list2
	}
	redPointer := list1
	bluePointer := list2
	head := list1
	for ; bluePointer != nil ; {
		if redPointer.Val > bluePointer.Val {
			tmp := bluePointer.Next
			bluePointer.Next = redPointer
			redPointer = bluePointer
			head = redPointer
			bluePointer = tmp
		} else {
			redTmp := redPointer.Next
			if redTmp == nil {
				redPointer.Next = bluePointer
				return head
			}
			for {
				if redTmp.Val > bluePointer.Val {
					b := bluePointer.Next
					redPointer.Next = bluePointer
					bluePointer.Next = redTmp
					bluePointer = b
					break
				} else {
					redPointer = redTmp
					break
				}
			}

		}
	}
	return head
}