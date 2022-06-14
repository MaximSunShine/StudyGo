package main

import (
	"fmt"
	"sync"
)

func main() {
	var a [5]int
	var A [5]int
	wg := sync.WaitGroup{}
	ch := make(chan int)
	for i := range a {
		wg.Add(2)
		go func(i int) {
			defer wg.Done()
			a[i] = i
			ch <- a[i]
		}(i)
		go func(i int) {
			defer wg.Done()
			A[i] = i * i
			ch <- A[i]
		}(i)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	sum := 0
	for i, _ := range A {
		sum += <-ch // почему 2 раза?
		sum += <-ch
		fmt.Printf("a = %v a*a = %v\n", a[i], A[i])
	}
	fmt.Print("Sum = ", sum)
}
