package main

import "fmt"

func canJump(nums []int) bool {
	var i, jumps, avaible int

	avaible = 1

	for i = 0; i < len(nums)-1 && i <= avaible; {
		if nums[i] == 0 {
			return false
		}
		jumps++
		if i+nums[i] >= len(nums)-1 {
			break
		}
		bestJ := 1

		for j := 1; j <= nums[i] && i+j < len(nums); j++ {
			if i+j+nums[i+j] > avaible {
				bestJ = j
				avaible = i + j + nums[i+j]
			}
		}
		/*if bestJ == 0 {
			return false
		}*/
		i = i + bestJ
	}
	return true
}
func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
}
