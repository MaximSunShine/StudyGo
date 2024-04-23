package main

import "fmt"

func removeDuplicates1(nums []int) int {
	if len(nums) < 3 {
		return len(nums)
	}
	i := 2
	for _, v := range nums[2:] {
		if v > nums[i-2] {
			nums[i] = v
			i++
		}
	}
	return i
}

func removeDuplicates(nums []int) int {
	if len(nums) < 3 {
		//fmt.Println(nums)
		return len(nums)
	}

	step := 1
	t := nums[0] == nums[1]

	for i := 2; i < len(nums); i++ { // 1111 112 123
		if nums[i] == nums[step] {
			if !t {
				step++
				nums[step] = nums[i]
				t = true
			}
		} else {
			step++
			nums[step] = nums[i]
			t = false
		}
	}
	//fmt.Println(nums[:step+1])
	return step + 1
}
func removeDuplicates2(nums []int) int {
	if len(nums) < 3 {
		//fmt.Println(nums)
		return len(nums)
	}

	step := 1
	t := nums[0] == nums[1]

	for i := 2; i < len(nums); i++ { // 1111 112 123
		if nums[i] == nums[step] {
			if !t {
				step++
				nums[step] = nums[i]
				t = true
			}
		} else {
			step++
			nums[step] = nums[i]
			t = false
		}
	}
	//fmt.Println(nums[:step+1])
	return step + 1
}
func main() {
	fmt.Println(removeDuplicates([]int{}))
	fmt.Println(removeDuplicates([]int{1}))
	fmt.Println(removeDuplicates([]int{1, 1}))
	fmt.Println(removeDuplicates([]int{1, 2}))
	fmt.Println(removeDuplicates([]int{1, 1, 1, 2, 2, 3}))
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 1, 2, 3, 3}))
}
