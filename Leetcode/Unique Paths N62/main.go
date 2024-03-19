package main

import (
	"fmt"
)

type Init struct {
	aa [][]int
}

var A = &Init{}

func (I *Init) exist() bool {
	if I.aa != nil {
		return true
	}
	return false
}

func (I *Init) initial() {
	I.aa = make([][]int, 100, 100)

	if I.aa[0] == nil {
		for i := 0; i < len(I.aa); i++ {
			I.aa[i] = append(I.aa[i], make([]int, 16, 16)...)
			for j := 0; j < len(I.aa[0]); j++ {
				if j == 0 || i == 0 {
					I.aa[i][j] = 1

					continue
				}
				if I.aa[i-1][j]+I.aa[i][j-1] < 0 {
					//fmt.Println(i, j)
					break
				}
				I.aa[i][j] = I.aa[i-1][j] + I.aa[i][j-1]

			}
		}
	}
	//fmt.Println(I.aa[len(I.aa)-1][len(I.aa)-1])
}
func uniquePathsBest(m int, n int) int {
	if m > n {
		return uniquePaths(n, m)
	}
	ret := 1
	for i := 0; i < m-1; i++ {
		ret *= (m + n - 2) - i
		ret /= i + 1
	}
	return ret
}

func uniquePaths(m int, n int) int {

	if !A.exist() {
		A.initial()
	}

	a := A.aa

	if m < n {
		m, n = n, m
	}

	fmt.Println(a[m-1][n-1])
	return a[m-1][n-1]

}
func main() {

	uniquePaths(3, 7)
	uniquePaths(17, 100)
	uniquePaths(17, 99)
	uniquePaths(17, 98)
	uniquePaths(17, 97)
	uniquePaths(17, 96)
	uniquePaths(17, 95)

	//uniquePaths(97, 16)

}
