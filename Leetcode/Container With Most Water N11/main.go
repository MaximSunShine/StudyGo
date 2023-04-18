package main

import (
	"fmt"
)

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxArea(a []int) int {
	l := 0
	r := len(a) - 1
	maxWater := 0
	for l < r {
		help := min(a[l], a[r])
		maxWater = max(maxWater, (r-l)*help)

		for ; a[l] <= help && l < r; l++ {
		}

		for ; a[r] <= help && l < r; r-- {
		}
	}
	return maxWater
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{1, 1}))
	fmt.Println(maxArea([]int{1, 2, 2, 2, 2, 18, 18, 2, 2, 2, 2, 1}))

}
