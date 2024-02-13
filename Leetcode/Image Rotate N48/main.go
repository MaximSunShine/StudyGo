package main

import "fmt"

func rotate(m [][]int) {
	/*for _, v := range m {
		fmt.Println(v)
	}*/
	for step := 0; step*2 <= len(m)-1; step++ {
		for i := step; i < len(m)-1-step; i++ {
			m[step][i], m[len(m)-1-i][step], m[len(m)-1-step][len(m)-1-i], m[i][len(m)-1-step] = m[len(m)-1-i][step], m[len(m)-1-step][len(m)-1-i], m[i][len(m)-1-step], m[step][i]
		}
	}

	/*fmt.Println()
	for _, v := range m {
		fmt.Println(v)
	}*/
}

func main() {
	rotate([][]int{})
	rotate([][]int{{1}})
	rotate([][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}})
	fmt.Println("hi")
}
