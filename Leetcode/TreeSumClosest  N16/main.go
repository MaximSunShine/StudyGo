package main

import (
	"fmt"
	"sort"
)

func doClose(x, y int) int {
	if x < y {
		x, y = y, x
	}
	return x - y
}

func threeSumClosest(nums []int, target int) int {

	sort.Ints(nums)

	closest := nums[0] + nums[1] + nums[2]
	if closest == target {
		return closest
	}
	length := len(nums)

	for i := 0; i < length-2; i++ {

		l := i + 1
		r := len(nums) - 1

		for l < r {

			help := nums[i] + nums[l] + nums[r]

			if help == target {
				return target
			}

			if doClose(help, target) < doClose(closest, target) {
				closest = help
			}

			if nums[l] == nums[r] {
				break
			}

			if help > target {
				for nums[r] == nums[r-1] {
					r--
				}
				if r > l {
					r--
				}
			} else {
				for nums[l] == nums[l+1] {
					l++
				}
				if l < r {
					l++
				}
			}
		}
		if nums[i] == nums[r] {
			break
		}
		for nums[i] == nums[i+1] {
			i++
		}
	}

	return closest
}
func main() {

	var nums = []int{1, 2, 10, 2, 4, -4}
	//var nums = []int{1, 2, 2, 2, 2, 2, 4}
	//var nums = []int{-9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//var nums = []int{1, 1, 1, 1, 1, 3, 3, 3, 3, 3, 6, 6, 6, 6, 6}
	//var nums = []int{0, 3, 97, 102, 200, 301}
	//var nums = []int{-1, 2, 1, -4}
	fmt.Println(threeSumClosest(nums, 9))

}
