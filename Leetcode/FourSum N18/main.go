package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	output := [][]int{}

	for l := 0; l < len(nums)-3; l++ {

		for r := len(nums) - 1; r > l+2; r-- {
			i := l + 1
			j := r - 1
			for i < j {
				if nums[l]+nums[i]+nums[j]+nums[r] == target {
					output = append(output, []int{nums[l], nums[i], nums[j], nums[r]})

					for nums[i] == nums[i+1] && i < j {
						i++
					}
					for nums[j] == nums[j-1] && j > i {
						j--
					}

					i++
					j--
				} else if nums[l]+nums[i]+nums[j]+nums[r] < target {
					i++
				} else {
					j--
				}

			}
			for nums[r] == nums[r-1] && r > l+2 {
				r--
			}
		}

		for nums[l] == nums[l+1] && l < len(nums)-2 {
			l++
		}
	}

	return output
}

func main() {

	//fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	fmt.Println(fourSum([]int{2, 2, 2, 2, 2}, 8))
}
