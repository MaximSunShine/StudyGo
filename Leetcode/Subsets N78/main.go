package main

import "fmt"

func subsets(nums []int) [][]int {
	n := len(nums)
	totalSubsets := 1 << n
	subsets := make([][]int, 0, totalSubsets)

	for i := 0; i < totalSubsets; i++ {
		currentSubset := make([]int, 0)
		for j := 0; j < n; j++ {
			if i&(1<<j) != 0 {
				currentSubset = append(currentSubset, nums[j])
			}
		}
		subsets = append(subsets, currentSubset)
	}

	return subsets
}

func subsets1(nums []int) [][]int {

	rez := [][]int{{}}

	for i := 0; i <= len(nums); i++ {
		rez = append(rez, combine(len(nums), i)...)
	}

	for i := 0; i < len(rez); i++ {
		for j := 0; j < len(rez[i]); j++ {
			rez[i][j] = nums[rez[i][j]-1]
		}
	}

	return rez
}
func fact(n int) int {
	rez := 1
	for i := 2; i <= n; i++ {
		rez *= i
	}
	return rez
}

func combine(n int, k int) [][]int {

	if k == 0 {
		return [][]int{}
	}

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
	fmt.Println(subsets([]int{-1, -2, -3}))
}
