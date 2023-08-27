package main

import "fmt"

func search(a []int, k int) int {

	if a[0] == k {
		return 0
	}
	if len(a) == 1 {
		if a[0] != k {
			return -1
		} else {
			return k
		}
	}

	l, t, m, r := 0, 0, 0, len(a)-1
	if k < a[0] {
		t = a[len(a)-1]
	} else {
		t = a[0]
	}

	for l != r {
		m = (r-l)/2 + l
		if a[m] > t {
			if l != m {
				l = m
			} else {
				l, m = r, r
				if a[r] == k {
					return r
				} else {
					break
				}
			}
		} else {
			r = m
		}
	}

	if a[len(a)-1] > a[0] {
		l = 0
		r = len(a) - 1
	} else if k < a[0] {
		l = r
		r = len(a) - 1
	} else {
		l = 0
		if r != 0 {
			r--
		}
	}

	for l != r {
		m = (r-l)/2 + l
		if a[m] < k {
			if l != m {
				l = m
			} else if a[r] == k {
				return r
			} else {
				break
			}
		} else {
			r = m
		}

	}

	return -1
}
func search1(nums []int, target int) int {

	ans := -1

	start, end := 0, len(nums)-1

	for start <= end {

		mid := start + (end-start)/2

		// fmt.Printf("Start: %v, Mid: %v, End: %v\n", start, mid, end)

		if nums[mid] == target {

			ans = mid
			break

			// left half is sorted
			// } else if nums[start] < nums[mid] { => Need to convert < to <= for 2nd sample tc, because when only 1 number remains and it does not match with target, none of the 3 if else statement matches

			/*

			   Uncomment print statement and run for following: [4,5,6,7,0,1,2], 3

			*/

		} else if nums[start] <= nums[mid] {

			if nums[start] <= target && target < nums[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}

			// right half is sorted
		} else if nums[mid] < nums[end] {

			if nums[mid] < target && target <= nums[end] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}

	return ans
}

func main() {

	fmt.Println(search([]int{5, 6, 7, 8, 9, 1, 2, 3, 4}, 8))
	fmt.Println(search([]int{5, 6, 7, 8, 9, 1, 2, 3, 4}, 3))
	fmt.Println(search([]int{5, 6, 7, 8, 9, 1, 2, 4}, 3))
	fmt.Println(search([]int{1}, 0))
	fmt.Println(search([]int{1, 3}, 3))
	fmt.Println(search([]int{3, 1}, 4))
}
