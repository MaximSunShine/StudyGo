package main

import "fmt"

func setZeroes(a [][]int) {
	x, y := len(a), len(a[0])
	zero_i, zero_j := false, false
	i, j := 0, 0

	for i = 0; i < x; i++ {
		if a[i][0] == 0 {
			zero_i = true
			break
		}
	}

	for j = 0; j < y; j++ {
		if a[0][j] == 0 {
			zero_j = true
			break
		}
	}

	for i = 1; i < x; i++ {
		for j = 1; j < len(a[0]); j++ {
			if a[i][j] == 0 {
				a[0][j] = 0
				a[i][0] = 0
			}
		}
	}

	for i = 1; i < x; i++ {
		if a[i][0] == 0 {
			for j = 1; j < y; j++ {
				a[i][j] = 0
			}
		}
	}

	for j = 1; j < y; j++ {
		if a[0][j] == 0 {
			for i = 1; i < x; i++ {
				a[i][j] = 0
			}
		}
	}

	if zero_i {
		for i = 0; i < x; i++ {
			a[i][0] = 0
		}
	}

	if zero_j {
		for j = 0; j < y; j++ {
			a[0][j] = 0
		}
	}

	fmt.Println(a)
}
func main() {
	setZeroes([][]int{{1}})
	setZeroes([][]int{{0}})
	setZeroes([][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}})
	setZeroes([][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}})

}
