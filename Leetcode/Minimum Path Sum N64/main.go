package main

import "fmt"

func minPathSum(a [][]int) int {
	n, m := len(a), len(a[0])

	if n*m == 1 {
		return a[0][0]
	}

	for i := 1; i < m; i++ {
		a[0][i] += a[0][i-1]
	}

	for i := 1; i < n; i++ {
		a[i][0] += a[i-1][0]
	}

	for i := 1; i < len(a); i++ {
		for j := 1; j < len(a[0]); j++ {
			if a[i-1][j] <= a[i][j-1] {
				a[i][j] += a[i-1][j]
			} else {
				a[i][j] += a[i][j-1]
			}
		}
	}
	return a[n-1][m-1]
}

func main() {
	fmt.Println(minPathSum([][]int{{3}}))
	fmt.Println(minPathSum([][]int{{4, 5}}))
	fmt.Println(minPathSum([][]int{{4}, {5}}))
	fmt.Println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
	fmt.Println(minPathSum([][]int{{1, 2, 3}, {4, 5, 6}}))

}
