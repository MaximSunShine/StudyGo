package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	p := head
	if p != nil {
		if p.Next != nil {
			head = head.Next
			p.Next, p.Next.Next = p.Next.Next, p
		}
		for p.Next != nil {
			if p.Next.Next != nil {
				p.Next, p.Next.Next, p.Next.Next.Next = p.Next.Next, p.Next.Next.Next, p.Next
				p = p.Next.Next

			} else {
				break
			}
		}
	}

	return head
}
func main() {
	var h, p *ListNode
	n := 1

	if n > 0 {
		p = &ListNode{1, nil}
		h = p
		for i := 2; i < n; i++ {
			p.Next = &ListNode{i, nil}
			p = p.Next
		}
	}

	p = swapPairs(h)
	for p != nil {
		fmt.Println(p.Val)
		p = p.Next
	}
}
