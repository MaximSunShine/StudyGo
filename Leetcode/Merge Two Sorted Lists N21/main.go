package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(x *ListNode, y *ListNode) *ListNode {

	if x == nil {
		return y
	} else if y == nil {
		return x
	}

	if x.Val > y.Val {
		x, y = y, x
	}
	first, help := x, x
	x = x.Next

	for x != nil && y != nil {

		if x.Val <= y.Val {
			help.Next = x
			x = x.Next
		} else {
			help.Next = y
			y = y.Next
		}

		help = help.Next
	}

	if x == nil {
		help.Next = y
	} else {
		help.Next = x
	}

	return first
}
func main() {

	a := []int{1, 3, 5, 7, 9}
	b := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26}

	x := &ListNode{}
	y := &ListNode{}

	h := x
	for i, v := range a {
		h.Val = v
		if i != len(a)-1 {
			h.Next = &ListNode{}
		}
		h = h.Next
	}

	h = y
	for i, v := range b {
		h.Val = v
		if i != len(b)-1 {
			h.Next = &ListNode{}
		}
		h = h.Next
	}

	for h := mergeTwoLists(x, y); ; {
		if h == nil {
			break
		}
		fmt.Println(h.Val)
		h = h.Next

	}

}
