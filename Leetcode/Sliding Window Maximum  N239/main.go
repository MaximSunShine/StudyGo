package main

import "fmt"

func maxSlidingWindow1(nums []int, k int) []int {
	result := make([]int, len(nums)-(k-1))
	queue := make([]int, 0, len(nums)-(k-1))
	for i, n := range nums {
		for len(queue) != 0 && n > nums[queue[len(queue)-1]] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
		if i >= k-1 {
			result[i-(k-1)] = nums[queue[0]]
			if queue[0] == i-(k-1) {
				queue = queue[1:]
			}
		}
	}
	return result
}

func findMax(nums []int, help int) (int, int) {
	max := nums[len(nums)-1]
	index := len(nums) - 1

	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == help {
			max = nums[i]
			index = i
			//fmt.Print("!")
			break
		}
		if max < nums[i] {
			max = nums[i]
			index = i
		}
	}

	return max, index
}

func maxSlidingWindow(nums []int, k int) []int {
	var a = make([]int, 0, len(nums)-k)
	l := 0
	r := l + k

	var max, index int
	max = 10000
	lens := len(nums)

	if r <= lens {
		max, index = findMax(nums[l:r], max)
		a = append(a, max)
		l++
		r++

		for r <= lens {
			if max <= nums[r-1] {
				max = nums[r-1]
				index = r - 1
				a = append(a, max)
				//fmt.Println("1 -> ", nums[r-1])

			} else if index >= l {
				a = append(a, max)
				//fmt.Println("2 -> ", max)

			} else {
				max, index = findMax(nums[l:r], max-1)
				index += l
				a = append(a, max)
				//fmt.Println("4 -> ", max)
			}

			l++
			r++
		}
	} else {

	}
	return a
}
func main() {

	//var a = []int{1, 3, -1, -3, 5, 3, 6, 7, 6, 5, 4, 2, 2, 1}
	//var a = []int{5, 4, 4,1,3,2 4, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5}
	//var a = []int{9, 8, 9, 8, 9, 8, 9, 8, 9, 8, 9, 8, 9, 8}
	var a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 9, 9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 2, 2, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5, 6, 6, 6, 6, 7, 7, 7, 7, 8, 8, 8, 8, 9, 9, 9, 9, 8, 8, 8, 8, 7, 7, 7, 7, 6, 6, 6, 6, 5, 5, 5, 5, 4, 4, 4, 4, 3, 3, 3, 3, 2, 2, 2, 2, 1, 1, 1, 1}
	k := 3

	fmt.Println(maxSlidingWindow1(a, k))
}
