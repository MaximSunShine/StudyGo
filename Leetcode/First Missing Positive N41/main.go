package main

import (
	"fmt"
	"sort"
)

func sortNew(nums []int) []int {

	for i := 0; i < len(nums); {
		if nums[i] < 1 {
			nums = append(nums[:i], nums[i+1:]...)
			continue
		}
		i++
	}
	return nums

}
func firstMissingPositive(a []int) int {

	for i := 0; i < len(a); i++ {
		for a[i] > 0 && a[i] <= len(a) && a[a[i]-1] != a[i] {
			a[i], a[a[i]-1] = a[a[i]-1], a[i]
		}
	}
	for i := 0; i < len(a); i++ {
		if a[i] != i+1 {
			return i + 1
		}
	}

	return len(a) + 1

}

func firstMissingPositive2(nums []int) int {

	nums = sortNew(nums)

	sort.Ints(nums)
	step := 1

	for i := 0; i < len(nums); i++ {
		if nums[i] < step {
			continue
		}
		if nums[i] == step {
			step++
		} else if nums[i] > step {
			return step
		}
	}

	return step
}
func firstMissingPositive1(a []int) int {
	n := len(a)

	for i := 0; i < n; i++ {
		for a[i] > 0 && a[i] <= n && a[i] != a[a[i]-1] {
			a[i], a[a[i]-1] = a[a[i]-1], a[i]
		}
	}

	for i := 0; i < n; i++ {
		if a[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}

func main() {

	fmt.Println(firstMissingPositive([]int{-1, 0, 1, 3}))       //2
	fmt.Println(firstMissingPositive([]int{-1, 0, 1, 1, 1, 3})) //2

	fmt.Println(firstMissingPositive([]int{-1, 0, 2}))        //1
	fmt.Println(firstMissingPositive([]int{1, 2, 0}))         //3
	fmt.Println(firstMissingPositive([]int{3, 4, -1, 1}))     //2
	fmt.Println(firstMissingPositive([]int{7, 8, 9, 11, 12})) //1
	fmt.Println(firstMissingPositive([]int{1}))               //2

	fmt.Println(firstMissingPositive([]int{-1, 4, 2, 1, 9, 10})) //3
}
