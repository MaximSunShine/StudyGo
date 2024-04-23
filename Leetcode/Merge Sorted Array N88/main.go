package main

import "fmt"

func merge(a []int, m int, b []int, n int) {
	m, n = m-1, n-1

	for i := n + m + 1; n >= 0; i-- {
		if m < 0 || b[n] > a[m] {
			a[i] = b[n]
			n--
		} else {
			a[i] = a[m]
			m--
		}
	}
}
func merge1(a []int, m int, b []int, n int) {

	a = a[:m]
	b = b[:n]

	for {
		if len(b) == 0 {
			return
		}
		if len(a) == 0 {
			a = append(a, b...)
			return
		}
		for len(a) > 0 && a[0] <= b[0] {
			a = a[1:]
		}
		if len(a) == 0 {
			continue
		}
		i := 0
		for i < len(b) && a[0] > b[i] {
			i++
		}
		a = append(a[0:0], append(append([]int{}, b[:i]...), a...)...)
		a = a[i:]
		b = b[i:]
	}
}

var aa = []int{4, 5, 6, 7, 0, 0, 0, 0}
var bb = []int{1, 2, 3, 8}

func main() {

	merge(aa, 4, bb, 4)
	fmt.Println(aa)

	aa = []int{1, 2, 3, 0, 0, 0}
	bb = []int{2, 5, 6}

	merge(aa, 3, bb, 3)

	fmt.Println(aa)

}
