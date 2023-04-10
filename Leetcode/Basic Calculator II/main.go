package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Example struct {
	value interface{}
	prev  *Example
	next  *Example
}

func doPlusMinus(queue *Example) *Example {
	start := queue

	for queue != nil {
		if queue.value == "+" || queue.value == "-" {
			//a, _ := strconv.Atoi(fmt.Sprintf("%d", queue.prev.value))
			//b, _ := strconv.Atoi(fmt.Sprintf("%d", queue.next.value.(int)))
			//queue.prev.value = reflect.ValueOf(queue.prev.value).Int() * reflect.ValueOf(queue.next.value).Int()
			if queue.value == "+" {
				queue.prev.value = int(queue.prev.value.(int)) + int(queue.next.value.(int))
			} else {
				queue.prev.value = queue.prev.value.(int) - queue.next.value.(int)
			}
			queue.prev.next = queue.next.next
			if queue.next.next != nil {
				queue.next.next.prev = queue.prev
			}
			queue = queue.next.next
		} else {
			queue = queue.next
		}
	}
	return start
}

func doMulDiv(queue *Example) *Example {

	start := queue

	for queue != nil {
		if queue.value == "*" || queue.value == "/" {
			//a, _ := strconv.Atoi(fmt.Sprintf("%d", queue.prev.value))
			//b, _ := strconv.Atoi(fmt.Sprintf("%d", queue.next.value.(int)))
			//queue.prev.value = reflect.ValueOf(queue.prev.value).Int() * reflect.ValueOf(queue.next.value).Int()
			if queue.value == "*" {
				queue.prev.value = int(queue.prev.value.(int)) * int(queue.next.value.(int))
			} else {
				queue.prev.value = queue.prev.value.(int) / queue.next.value.(int)
			}
			queue.prev.next = queue.next.next
			if queue.next.next != nil {
				queue.next.next.prev = queue.prev
			}
			queue = queue.next.next
		} else {
			queue = queue.next
		}
	}
	return start
}

func readNextValue(s string) (string, string) {
	if string(s[0]) == "+" || string(s[0]) == "-" || string(s[0]) == "*" || string(s[0]) == "/" {
		return string(s[0]), s[1:]
	}

	num := ""
	for i, v := range s {
		val := string(v)
		if val == "+" || val == "-" || val == "*" || val == "/" {
			return num, s[i:]
		} else {
			num += val
		}

	}
	return num, ""
}

func convertToQueue(s string, queue *Example) *Example {

	var val interface{}
	v := ""

	if len(s) > 0 {
		v, s = readNextValue(s)
		v = strings.Trim(v, " ")
		n, err := strconv.Atoi(v)
		if err == nil {
			val = n
		} else {
			val = v
		}
		queue = new(Example)
		queue.value = val
	}

	start := queue

	for len(s) > 0 {
		v, s = readNextValue(s)
		v = strings.Trim(v, " ")
		n, err := strconv.Atoi(v)
		if err == nil {
			val = n
		} else {
			val = v
		}
		queue.next = new(Example)
		queue, queue.next.prev = queue.next, queue
		queue.value = val
	}
	return start
}

func calculate(s string) int {
	var start *Example
	start = convertToQueue(s, start)
	start = doMulDiv(start)
	start = doPlusMinus(start)
	return start.value.(int)
}
func main() {
	fmt.Println(calculate("2*3*4"))

}
