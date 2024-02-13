package main

import "fmt"

func jump(nums []int) int {

	if len(nums) == 1 {
		return 0
	}
	if len(nums) == 2 {
		return 1
	}

	step := 0
	n := 0
	for n < len(nums)-1 {
		l := nums[n]

		if n+l >= len(nums)-1 {
			return step + 1
		}

		max := nums[n+1]
		imax := n + 1
		k := 0
		for i := n + 2; i < len(nums)-1 && i <= n+l; i++ {
			k++
			if max <= nums[i]+k {
				max = nums[i]
				imax = i
				k = 0
			}
		}
		step++
		n = imax

	}

	return step

}
func main() {
	fmt.Println(jump([]int{7, 0, 9, 6, 9, 6, 1, 7, 9, 0, 1, 2, 9, 0, 3}))
	fmt.Println(jump([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 0}))
	fmt.Println(jump([]int{1, 2, 1, 1, 1}))
	fmt.Println(jump([]int{1, 1, 1, 1, 1}))
	fmt.Println(jump([]int{1, 0}))
	fmt.Println(jump([]int{2}))
	fmt.Println(jump([]int{4, 1, 1, 3, 1, 1, 1}))
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
}
