package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	l, r := &ListNode{101, head}, head.Next
	head = l

	for r != nil {
		if l.Next.Val == r.Val {
			for r != nil && l.Next.Val == r.Val {
				r = r.Next
			}
			l.Next = r
			if r != nil {
				r = r.Next
			}
		} else {
			r = r.Next
			l = l.Next
		}
	}

	return head.Next
}

func deleteDuplicates1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	l, r := head, head.Next

	for r != nil {
		if l == head {
			if l.Val == r.Val {
				for r != nil && l.Val == r.Val {
					r = r.Next
				}
				head, l = r, r
				if r != nil {
					r = r.Next
				}
			} else {
				r = r.Next
				break
			}
		}
	}

	for r != nil {
		if l.Next.Val == r.Val {
			for r != nil && l.Next.Val == r.Val {
				r = r.Next
			}
			l.Next = r
			if r != nil {
				r = r.Next
			}
		} else {
			r = r.Next
			l = l.Next
		}
	}

	return head
}
func main() {

	a := []int{1, 2, 3, 3, 4, 4, 5}
	a = []int{1, 1, 1, 2, 3}
	a = []int{1, 1}
	a = []int{1, 1, 2, 3, 3}
	a = []int{1, 2, 2}
	head := &ListNode{Val: a[0]}
	help := head
	for _, v := range a[1:] {
		help.Next = &ListNode{v, nil}
		help = help.Next
	}

	deleteDuplicates(head)

	fmt.Println("Done")

}
