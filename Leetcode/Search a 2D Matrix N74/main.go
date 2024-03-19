package main

import "fmt"

func searchMatrix1(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])

	left, right := 0, m*n-1
	for left <= right {
		mid := left + (right-left)/2
		midValue := matrix[mid/n][mid%n]
		if midValue == target {
			return true
		} else if midValue < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}
func searchMatrix(a [][]int, target int) bool {

	down, up, l, r, mid := 0, len(a)-1, 0, len(a[0])-1, 0

	//find := false

	for down < up {
		mid = down + (up-down)/2 + 1
		if a[mid][0] == target {
			return true
		}
		if a[mid][0] < target {
			down = mid
		} else {
			up = mid - 1
		}
	}

	for l < r {
		mid = l + (r-l)/2
		if a[down][mid] == target {
			return true
		}
		if a[down][mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return a[down][l] == target
}

func main() {
	fmt.Println(searchMatrix([][]int{{1}}, 2))
	fmt.Println(searchMatrix([][]int{{1}}, 1))
	fmt.Println(searchMatrix([][]int{{1}}, 0))
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3))
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13))

}
