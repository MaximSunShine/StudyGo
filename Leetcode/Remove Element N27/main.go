package main

import "fmt"

func removeElement(nums []int, val int) int {
	l := 0
	for i, num := range nums {
		if num != val { //&&
			if i != l {
				nums[l] = num
			}
			l++
		}
	}
	return l
}

func main() {
	fmt.Println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
	fmt.Println(removeElement([]int{}, 1))
	fmt.Println(removeElement([]int{1}, 1))

	fmt.Println(removeElement([]int{3, 3}, 3))

	fmt.Println(removeElement([]int{1}, 2))
	fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))
}
