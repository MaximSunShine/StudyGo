package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func trap(a []int) int {

	l, r, sum, high, i := 0, len(a)-1, 0, 0, 0

	for l < r-1 {
		high = max(min(a[l], a[r]), high)
		if a[l] <= a[r] {
			l++
			i = l
		} else {
			r--
			i = r
		}
		if a[i] < high {
			sum += high - a[i]
		}
	}
	return sum
}

func trap2(a []int) int {

	l, r, sum, max, i := 0, len(a)-1, 0, 0, 0

	for l < r-1 {

		if a[l] <= a[r] {
			if max < a[l] {
				max = a[l]
			}
			l++
			i = l
		} else {
			if max < a[r] {
				max = a[r]
			}
			r--
			i = r
		}
		if a[i] < max {
			sum += max - a[i]
		}
	}

	return sum
}
func trap1(a []int) int {
	l := 0
	r := len(a) - 1
	sum := 0

	if r == 0 {
		return 0
	}
	step := 0
	i := 0
	for l < r {
		for l < r && (a[l] < a[l+1] || a[l] <= step || a[r] < a[r-1] || a[r] <= step) {
			if a[l] < a[l+1] || a[l] <= step {
				l++
			}

			if a[r] < a[r-1] || a[r] <= step {
				r--
			}
		}
		if r-l <= 1 {
			break
		}

		if a[l] > a[r] {
			step = a[r]
		} else {
			step = a[l]
		}

		if l < r-1 {
			for k := l + 1; k <= r-1; k++ {
				if a[k] < step {
					if a[k] < i {
						sum += step - i
					} else {
						sum += step - a[k]
					}

				}
			}
		}
		if a[l] < a[r] {
			l++
		} else if a[r] < a[l] {
			r--
		} else {
			l++
			r--
		}
		i = step
	}

	return sum
}

func main() {

	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})) // 6
	fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))                   // 9
}
