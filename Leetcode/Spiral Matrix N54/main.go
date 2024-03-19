package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	i := len(matrix[0])
	j := len(matrix)
	l := 0
	n := i * j
	a := make([]int, 0, n)

	for {

		a = append(a, matrix[l][l:i]...) //->
		for k := l + 1; k < j; k++ {
			a = append(a, matrix[k][i-1]) //!
		}

		for k := i - 2; k >= l && (j-1 != l); k-- { //<-
			a = append(a, matrix[j-1][k])
		}

		for k := j - 2; k > l && (i-1 != l); k-- { //1
			a = append(a, matrix[k][l])
		}

		l++
		i--
		j--

		if len(a) == n {
			break
		}
		/*if l >= i || l >= j {
			break
		}*/
	}

	return a
}
func main() {
	fmt.Println(spiralOrder([][]int{{1}}))
	fmt.Println(spiralOrder([][]int{{1, 2, 3, 4, 5, 6, 7}}))
	fmt.Println(spiralOrder([][]int{{1}, {2}, {3}, {4}, {5}, {6}, {7}}))
	fmt.Println(spiralOrder([][]int{ //1 2 3 6 9 8 7 4 5
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9}}))

	fmt.Println(spiralOrder([][]int{ //1 2 3 4 8 12 11 10 9 5 6 7
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12}}))

	fmt.Println(spiralOrder([][]int{ //1 2 3 6 9 12 11 10 7 4 5 8
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{10, 11, 12}}))

	fmt.Println(spiralOrder([][]int{
		//1 2 3 4 5 10 15 20 19 18 17 16 11 6 7 8 9 14 13 12
		//1 2 3 4 5 10 15 20 19 18 17 16 11 6 7 8 9 14 13 12
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20}}))

	fmt.Println(spiralOrder([][]int{
		// 1 2 3 4 8 12 16 20 19 18 17 13 9 5 6 7 11 15 14 10
		//[1 2 3 4 8 12 16 20 19 18 17 13 9 5 6 7 11 15 10]
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
		{17, 18, 19, 20}}))
	//fmt.Println("Hi")
}
