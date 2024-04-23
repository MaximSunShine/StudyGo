package main

import (
	"fmt"
)

func fact(n int) int {
	rez := 1
	for i := 2; i <= n; i++ {
		rez *= i
	}
	return rez
}

func combine(n int, k int) [][]int {

	a := make([]int, k, k)
	for i := 0; i < k; i++ {
		a[i] = n - k + i + 1
	}
	var help = make([]int, k, k)
	copy(help, a)
	rez := make([][]int, 0, fact(n)/(fact(k)*fact(n-k)))
	rez = append(rez, help)

	cur := 0

	for a[k-1] > k {
		for a[cur] > cur+1 { //decrease first position
			a[cur]-- // decrease first position
			var help = make([]int, k, k)
			copy(help, a)
			rez = append(rez, help)
		}
		for cur < k-1 { // move right
			cur++                // move right
			if a[cur] == cur+1 { //
				continue
			} else {
				a[cur]--
				cur--
				for ; cur >= 0; cur-- {
					a[cur] = a[cur+1] - 1
				}
				cur++
				var help = make([]int, k, k)
				copy(help, a)
				rez = append(rez, help)

				break
			}
		}
	}

	return rez
}
func main() {

	fmt.Println(combine(4, 3))
	fmt.Println(combine(4, 2))
}
