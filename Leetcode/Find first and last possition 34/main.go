package main

import "fmt"

func searchRange(nums []int, target int) []int {
	l, r := 0, len(nums)-1
	exist := false
	for l <= r {
		m := r - (r-l)/2
		if target == nums[m] {
			exist = true
		}
		if target <= nums[m] {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	if !exist {
		return []int{-1, -1}
	}
	lidx := l

	l, r = lidx, len(nums)-1
	for l <= r {
		m := r - (r-l)/2
		if target >= nums[m] {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	ridx := r
	return []int{lidx, ridx}
}

func searchRange1(nums []int, target int) []int {

	fs, ls := -1, -1

	l, r := 0, len(nums)-1

	if len(nums) == 0 {
		return []int{-1, -1}
	}

	if nums[l] == target && nums[r] == target {
		return []int{l, r}
	}

	help := 0

	for r-l > 0 {
		if nums[l] != target || nums[r] != target {
			if nums[l] < target {
				help = l + (r-l)/2
			}

			if nums[help] < target && help != l {
				l = help
			} else if nums[help] > target && help != r {
				r = help
			} else {
				if nums[l] < target {
					l++
				}
				if nums[r] > target {
					r--
				}
			}
		} else {
			break
		}

		if r-l == 1 && nums[l] != target && nums[r] != target { //если цели не существует вообще
			break
		}
	}

	if nums[l] == target {
		fs = l
		ls = r

	}

	return []int{fs, ls}
}

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 10}, 8))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	fmt.Println(searchRange([]int{}, 0))
	fmt.Println(searchRange([]int{1, 4}, 4))
}
