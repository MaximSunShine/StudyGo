package main

import "fmt"

func search(nums []int, target int) bool {
	l, r, mid := 0, len(nums)-1, 0

	if nums[0] == target || nums[len(nums)-1] == target {
		return true
	}

	if nums[len(nums)-1] <= nums[0] {
		for l < r {
			if l < r-1 && nums[l] == nums[r] {
				for ; l < r && nums[l] == nums[r]; r-- {
				}
				if l < r {
					return search(nums[l:r+1], target)
				}
			}
			mid = l + (r-l)/2
			if nums[mid] <= nums[r] /*&& nums[r] < nums[l]*/ {
				r = mid
			} else {
				l = mid + 1
			}
		}
	}
	if r == 0 /*|| r == len(nums)-1 */ {
		return false
	}

	if target < nums[0] {
		r = len(nums) - 1
	} else {
		l = 0
		if r > 0 {
			r--
		}
	}

	if nums[l] == target || nums[r] == target {
		return true
	}

	for l < r {
		mid = l + (r-l)/2
		if nums[mid] == target {
			return true
		}
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return false
}
func search2(nums []int, target int) bool {
	l, r, mid := 0, len(nums)-1, 0

	if nums[0] == target || nums[len(nums)-1] == target {
		return true
	}

	if nums[len(nums)-1] <= nums[0] {
		for l < r {
			for ; nums[l] == nums[r]; l, r = l+1, r-1 {
			}
			mid = l + (r-l)/2
			if nums[mid] < nums[r] && nums[r] < nums[l] {
				r = mid
			} else {
				l = mid + 1
			}
		}
	}

	if r == 0 || r == len(nums)-1 {
		return false
	}

	if target < nums[0] {
		r = len(nums) - 1
	} else {
		l = 0
	}

	if nums[l] == target || nums[r] == target {
		return true
	}

	for l < r {
		mid = l + (r-l)/2
		if nums[mid] == target {
			return true
		}
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return false
}
func main() {

	fmt.Println(search([]int{1, 2, 3, 0, 0, 0}, 2) == true)
	fmt.Println(search([]int{2, 2, 3, 3, 0, 1, 1}, 0) == true)
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 5) == true)
	fmt.Println(search([]int{1, 3, 5}, 3) == true)
	fmt.Println(search([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1}, 2) == true)
	fmt.Println(search([]int{8, 1, 2, 2, 5, 8}, 5) == true)
	fmt.Println(search([]int{1, 1, 2, 2, 5, 8}, 0) == false)
	fmt.Println(search([]int{2, 2, 2, 2, 2, 2, 2, 2}, 0) == false)
	fmt.Println(search([]int{1, 0, 1, 1, 1}, 0) == true)
	fmt.Println(search([]int{2}, 0) == false)
	fmt.Println(search([]int{2}, 2) == true)
	fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 0) == true)
	fmt.Println(search([]int{2, 5, 6, 0, 0, 0, 1, 2}, 3) == false)
	fmt.Println(search([]int{2, 3, 5, 6, 0, 0, 1, 2}, 0) == true)
	fmt.Println(search([]int{2, 3, 5, 6, 0, 0, 0, 1, 2}, 3) == true)

}
