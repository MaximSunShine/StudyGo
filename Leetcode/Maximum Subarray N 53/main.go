package main

import (
	"fmt"
)

func maxSubArray2(nums []int) int {
	pre := 0
	res := nums[0]
	for _, num := range nums {
		pre = max(pre+num, num)
		res = max(res, pre)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSubArraymy(nums []int) int {
	max := nums[0]
	here := 0
	if nums[0] > 0 {
		here = nums[0]
	}

	for _, n := range nums[1:] {
		here += n
		if here > max {
			max = here
		}
		if here < 0 {
			here = 0
		}
	}
	return max
}

func maxSubArray1(a []int) int {

	sum, max := a[0], a[0]
	if a[0] < 0 {
		sum = 0
	}
	for i := 1; i < len(a); i++ {
		if a[i]+sum <= 0 {
			sum = 0
			if a[i] > max {
				max = a[i]
			}
		} else if a[i] > 0 {
			sum += a[i]
			if max < sum {
				max = sum
			}
		} else if a[i] < 0 {
			sum += a[i]
		}

	}

	return max
}
func maxSubArray(nums []int) int {
	max, sum := nums[0], nums[0]
	for _, v := range nums[1:] {
		if sum < 0 {
			sum = v
		} else {
			sum += v
		}
		if max < sum {
			max = sum
		}
	}
	return max
}
func main() {
	fmt.Println(maxSubArray([]int{-3, -2, -10, -5, -1, -1, 0}))   //1
	fmt.Println(maxSubArray([]int{-1, 1, 2, 1}))                  //4
	fmt.Println(maxSubArray([]int{2, 1, -3, 4, -1, 2, 1, -5, 4})) //6
	fmt.Println(maxSubArray([]int{-1}))                           //-1
	fmt.Println(maxSubArray([]int{1}))                            //1
	fmt.Println(maxSubArray([]int{5, 4, -1, 7, 8}))               //23
	fmt.Println(maxSubArray([]int{1, 1, 1}))                      //3
}
