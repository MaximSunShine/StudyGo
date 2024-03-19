package main

import "fmt"

func sortColors1(nums []int) {

	const RED, BLUE = 0, 2

	i, idx_red, idx_blue := 0, 0, len(nums)-1

	for ; i <= idx_blue; i += 1 {

		if nums[i] == RED {
			nums[i], nums[idx_red] = nums[idx_red], nums[i]
			idx_red += 1

		} else if nums[i] == BLUE {
			nums[i], nums[idx_blue] = nums[idx_blue], nums[i]
			idx_blue -= 1

			// i stay here for one more check on next iteration
			i -= 1

		}
	}
	return
}

func sortColors(a []int) {
	var zero, one, two []int
	for _, v := range a {
		switch v {
		case 0:
			zero = append(zero, v)
		case 1:
			one = append(one, v)
		case 2:
			two = append(two, v)
		}
	}

	copy(a, append(zero, append(one, two...)...))
}

func main() {
	a := []int{2, 0, 2, 1, 1, 0}
	sortColors(a)
	fmt.Println(a)
	a = []int{2, 0, 1}
	sortColors(a)
	fmt.Println(a)
}
