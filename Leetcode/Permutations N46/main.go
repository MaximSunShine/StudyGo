package main

import (
	"fmt"
	"sort"
	"sync"
)

func find(a []int, aa *[][]int, bottom int) {

	for i := 0; i < len(a)-bottom; i++ {
		if len(a)-bottom == 2 {
			b, c := make([]int, len(a), len(a)), make([]int, len(a), len(a))
			copy(b, a)
			*aa = append(*aa, b[:])
			a[bottom], a[bottom+1] = a[bottom+1], a[bottom]
			copy(c, a)
			*aa = append(*aa, c[:])

			return
		} else {
			find(a, aa, bottom+1)

			if len(a)-bottom < 4 {
				a[bottom], a[bottom+1] = a[bottom+1], a[bottom]
			}
			if len(a)-bottom > 3 {

				/*help := a[bottom]
				a = append(a[:bottom], a[bottom+1:]...)
				a = append(a, help)*/

				for i := bottom; i < len(a)-1; i++ {
					if a[i] == a[i+1] {
						i++
						continue
					}
					a[i], a[i+1] = a[i+1], a[i]
				}
			}
		}
	}
	return
}

func permute2(a []int) [][]int {
	if len(a) < 2 {
		return [][]int{a}
	}

	N := 1
	for i := 2; i <= len(a); i++ {
		N *= i
	}

	aa := make([][]int, 0, N)

	find(a, &aa, 0)

	fmt.Println(len(aa))
	return aa

}
func find1(a []int, ch chan []int, wg *sync.WaitGroup, bottom int) {

	for i := 0; i < len(a)-bottom; i++ {
		if len(a)-bottom == 2 {
			b, c := make([]int, len(a), len(a)), make([]int, len(a), len(a))
			copy(b, a)
			ch <- b
			a[bottom], a[bottom+1] = a[bottom+1], a[bottom]
			copy(c, a)
			ch <- c

			wg.Done()
			return
		} else {
			wg.Add(1)
			find1(a, ch, wg, bottom+1)

			if len(a)-bottom < 4 {
				a[bottom], a[bottom+1] = a[bottom+1], a[bottom]
			}
			if len(a)-bottom > 3 {
				for i := bottom; i < len(a)-1; i++ {
					a[i], a[i+1] = a[i+1], a[i]
				}
			}

		}

	}
	wg.Done()
	return
}

func permute1(a []int) [][]int {
	if len(a) < 2 {
		return [][]int{a}
	}

	ch := make(chan []int)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go find1(a, ch, wg, 0)

	go func() {
		wg.Wait()
		close(ch)
	}()

	aa := [][]int{}
	for c := range ch {
		aa = append(aa, c)
	}

	//fmt.Println(len(aa))
	return aa

}
func permute(a []int) [][]int {
	aa := [][]int{}
	rez := make([]int, 0, len(a))

	var f func([]int)

	f = func(o []int) {
		if len(o) == 0 {
			aa = append(aa, append([]int{}, rez[:]...))
			return
		}
		for i := len(o) - 1; i >= 0; i-- {
			if i > 0 && o[i] == o[i-1] {
				continue
			}
			rez = append(rez, o[i])
			help := append([]int{}, o[:i]...)
			help = append(help, o[i+1:]...)
			f(help)
			rez = rez[:len(rez)-1]

		}
	}

	f(a[:])
	fmt.Println(len(aa))

	return aa
}
func permuteUnique(nums []int) [][]int {
	var result [][]int
	helper(nums, 0, len(nums), &result)
	fmt.Println(len(result))
	return result
}
func helper(nums []int, l, n int, result *[][]int) {
	if l == n {
		arr := make([]int, n)
		copy(arr, nums)
		*result = append(*result, arr)
	}
	//hash := make(map[int]bool)
	for i := l; i < n; i++ {
		//if !hash[nums[i]] {
		//hash[nums[i]] = true
		nums[i], nums[l] = nums[l], nums[i]
		helper(nums, l+1, n, result)
		nums[i], nums[l] = nums[l], nums[i]
		//}
	}
}
func main() {
	fmt.Println(permuteUnique([]int{1, 2, 3, 4, 5}))

	//fmt.Println(permute2([]int{1, 1, 3}))
	n := []int{0, 0, 1, 1, 2}
	sort.Ints(n)

	fmt.Println(permute(n))
}
