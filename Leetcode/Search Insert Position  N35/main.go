package main

import "fmt"

func searchInsert(nums []int, target int) int {

	l, r := 0, len(nums)-1
	s := 0

	if nums[l] >= target {
		return 0
	}

	if nums[r] == target {
		return r
	}

	if nums[r] < target {
		return r + 1
	}

	for l <= r {
		s = (r + l + 1) / 2
		if nums[s] < target {
			l = s + 1
		} else if nums[s] > target {
			r = s - 1
		} else {
			return s
		}

	}

	return l
}

func main() {

	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5)) //2

	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2)) //1

	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7)) //4

	fmt.Println("0 - ", searchInsert([]int{1, 3, 5, 6, 7, 8, 11, 23, 24, 35}, 0))   //0
	fmt.Println("0 - ", searchInsert([]int{1, 3, 5, 6, 7, 8, 11, 23, 24, 35}, 1))   //0
	fmt.Println("1 - ", searchInsert([]int{1, 3, 5, 6, 7, 8, 11, 23, 24, 35}, 2))   //1
	fmt.Println("1 - ", searchInsert([]int{1, 3, 5, 6, 7, 8, 11, 23, 24, 35}, 3))   //1
	fmt.Println("2 - ", searchInsert([]int{1, 3, 5, 6, 7, 8, 11, 23, 24, 35}, 4))   //2
	fmt.Println("4 - ", searchInsert([]int{1, 3, 5, 6, 7, 8, 11, 23, 24, 35}, 7))   //4
	fmt.Println("5 - ", searchInsert([]int{1, 3, 5, 6, 7, 8, 11, 23, 24, 35}, 8))   //5
	fmt.Println("6 - ", searchInsert([]int{1, 3, 5, 6, 7, 8, 11, 23, 24, 35}, 9))   //6
	fmt.Println("9 - ", searchInsert([]int{1, 3, 5, 6, 7, 8, 11, 23, 24, 35}, 35))  //8
	fmt.Println("10 - ", searchInsert([]int{1, 3, 5, 6, 7, 8, 11, 23, 24, 35}, 36)) //9

}
