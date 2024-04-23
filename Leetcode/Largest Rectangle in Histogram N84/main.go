package main

import "fmt"

func largestRectangleArea(a []int) (res int) {
	a = append(a, -1)
	n := len(a)
	l := make([]int, n)
	// l[i] - index of the first <= a[i] on the left
	// if a[i-1] <= a[i]: l[i] = i-1
	// else for  (...a[l[i-1]], a[i-1]) > a[i]: S = a[j]*(i-l[j]-1)
	// and l[i] is the first j: a[j] <= a[i]
	for i := 0; i < n; i++ {
		j := i - 1
		for ; j > -1 && a[j] > a[i]; j = l[j] {
			res = max(a[j]*(i-l[j]-1), res)
		}
		l[i] = j
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func largestRectangleArea1(a []int) int {

	if len(a) == 1 {
		return a[0]
	}
	mainMin := a[0]

	min := 100000
	help := 0
	max := a[0]
	for i := 0; i < len(a); i++ {
		if mainMin > a[i] {
			mainMin = a[i]
		}
		for i < len(a)-1 && a[i] <= a[i+1] {
			i++
		}
		min = a[i]
		for j := i; j >= 0; j-- {
			if a[j] == mainMin {
				help = (i + 1) * a[j]
				j = 0
			} else {
				if min > a[j] {
					min = a[j]
				}
				help = (i - j + 1) * min
			}

			if max < help {
				max = help
			}
		}
	}

	return max
}

func main() {

	fmt.Println(largestRectangleArea([]int{2, 4}) == 4)
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}) == 10)

}
