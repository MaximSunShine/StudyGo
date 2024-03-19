package main

import "fmt"

func generateMatrix(n int) [][]int {

	aa := make([][]int, n, n)
	for i, _ := range aa {
		aa[i] = append(aa[i], make([]int, n, n)...)
	}

	l := 0
	step := 1
	i := n
	j := n

	for {
		//aa[l][l:i] = step
		//step++
		//a = append(a, matrix[l][l:i]...) //->
		for k := l; k < i; k++ {
			aa[l][k] = step
			step++
			//a = append(a, aa[k][i-1]) //!
		}

		for k := l + 1; k < j; k++ {
			aa[k][i-1] = step
			step++
			//a = append(a, aa[k][i-1]) //!
		}
		for k := i - 2; k >= l && j-1 != l; k-- { //<-
			aa[j-1][k] = step
			step++
			//a = append(a, aa[j-1][k])
		}
		for k := j - 2; k > l && i-1 != l; k-- { //1
			aa[k][l] = step
			step++
			//a = append(a, aa[k][l])
		}

		l++
		i--
		j--

		if l >= i || l >= j {
			break
		}
	}

	return aa
}
func main() {

	fmt.Println(generateMatrix(1))
	fmt.Println(generateMatrix(2))
	fmt.Println(generateMatrix(3))
	fmt.Println(generateMatrix(4))
	fmt.Println(generateMatrix(20))

}
