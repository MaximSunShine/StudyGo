package main

import "fmt"

// * Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if k < 1 || head == nil {
		return head
	}

	cur := head
	lenList := 0
	a := make([]*ListNode, 0, 500)
	for ; cur.Next != nil; lenList++ {
		a = append(a, cur)
		cur = cur.Next
	}

	lenList++
	realMove := k % lenList
	if realMove == 0 {
		return head
	}

	a = append(a, cur)
	cur.Next = head

	cur = a[len(a)-1-realMove]
	head = cur.Next
	cur.Next = nil

	return head

}

func createList(k int) *ListNode {

	var first *ListNode

	first = &ListNode{1, nil}
	cur := first

	for i := 2; i <= k; i++ {
		cur.Next = &ListNode{i, nil}
		cur = cur.Next
	}
	return first
}

func main() {
	fmt.Println(rotateRight(createList(1), 1))
	fmt.Println(rotateRight(createList(5), 2))
	fmt.Println(rotateRight(createList(3), 1000))

}
