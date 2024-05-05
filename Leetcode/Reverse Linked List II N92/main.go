package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, l int, r int) *ListNode {
	if head.Next == nil || l == r {
		return head
	}

	var first, last *ListNode
	help := &ListNode{-1, head}
	head = help
	for {
		r--
		l--
		if l == 0 {
			first = help
		}
		if r == 0 {
			last = help.Next
			break
		}
		help = help.Next
	}

	for first.Next != last {
		help = first.Next
		first.Next = first.Next.Next

		help.Next = last.Next
		last.Next = help
	}

	return head.Next
}

func main() {
	reverseBetween(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}, 1, 4)

}
