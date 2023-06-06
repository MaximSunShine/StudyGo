package main

import (
	"fmt"
)

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	n++
	p := head
	del := head
	for p != nil {
		if n != 0 {
			n--
		} else {
			del = del.Next
		}
		p = p.Next
	}

	if n == 1 {
		head = head.Next
	} else {
		del.Next = del.Next.Next
	}
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	start := new(ListNode)
	start.Val = 0

	p := start

	for i := 1; i < 10; i++ {
		p.Next = new(ListNode)
		p = p.Next
		p.Val = i
	}

	p = removeNthFromEnd(start, 10)

	for p != nil {
		fmt.Print(p.Val, " ")
		p = p.Next
	}

}
