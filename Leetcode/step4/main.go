package main

import (
	"container/list"
	"fmt"
)

func Push(elem interface{}, queue *list.List) {
	queue.PushBack(elem)
}

func Pop(queue *list.List) interface{} {
	elem := queue.Front()
	if queue.Len() > 0 {
		queue.Remove(elem)
	}
	return elem

}

func printQueue(queue *list.List) {
	for elem := queue.Front(); elem != nil; elem = elem.Next() {
		fmt.Print(elem.Value)
	}
}

func ReverseList(l *list.List) *list.List {
	r := new(list.List)
	for l.Len() > 0 {
		r.PushFront(l.Remove(l.Front()))
	}
	return r
}

func main() {
	mylist := list.New()
	mylist.PushBack(0)
	mylist.PushBack(1)
	mylist.PushBack(2)
	mylist.PushBack(3)

	Push(4, mylist)
	Pop(mylist)
	printQueue(mylist)
	fmt.Println()
	printQueue(ReverseList(mylist))

}
