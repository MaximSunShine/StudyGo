package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	var moreHead = &ListNode{-1, nil}
	var moreEnd = moreHead

	cur := head
	prevCur := cur
	for cur != nil {
		first := cur
		last := cur
		for cur != nil && cur.Val >= x {
			last = cur
			cur = cur.Next
		}

		if first == cur {
			prevCur = cur
			cur = cur.Next
		} else {
			moreEnd.Next = first
			moreEnd = last

			if first == head {
				if cur == nil {
					return head
				}
				head = cur
			} else {
				prevCur.Next = cur
			}
		}
	}

	prevCur.Next = moreHead.Next
	moreEnd.Next = nil

	return head
}

func main() {
	head := &ListNode{1, &ListNode{1, nil}}
	head = partition(head, 0)
	for head != nil {
		fmt.Print(head.Val, ";")
		head = head.Next
	}

	fmt.Println()

	head = &ListNode{2, &ListNode{1, nil}}
	head = partition(head, 2)
	for head != nil {
		fmt.Print(head.Val, ";")
		head = head.Next
	}

	fmt.Println()

	head = &ListNode{1, &ListNode{4, &ListNode{3, &ListNode{2, &ListNode{5, &ListNode{2, nil}}}}}}
	head = partition(head, 3)
	for head != nil {
		fmt.Print(head.Val, ";")
		head = head.Next
	}

}
