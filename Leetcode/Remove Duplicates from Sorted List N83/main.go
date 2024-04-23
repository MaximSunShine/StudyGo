package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	help := head

	for help.Next != nil {
		if help.Val == help.Next.Val {
			help.Next = help.Next.Next
			continue
		}
		help = help.Next
	}

	return head
}

func main() {
	a := []int{1, 2, 3, 3, 4, 4, 5}
	a = []int{1, 1, 1, 2, 3}
	a = []int{1, 1}
	a = []int{1, 1, 2}
	a = []int{1, 1, 2, 3, 3}
	head := &ListNode{Val: a[0]}
	help := head
	for _, v := range a[1:] {
		help.Next = &ListNode{v, nil}
		help = help.Next
	}

	deleteDuplicates(head)
}
