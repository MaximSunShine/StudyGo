package main

import "fmt"

func maximalRectangle1(a [][]byte) int {

	var sum byte = 0
	jj, ii := len(a[0]), len(a)

	for j := 0; j < jj; j++ {
		sum = 0
		for i := 0; i < ii; i++ {
			sum = (a[i][j] - 48) * (sum + 1)
			a[i][j] = sum
		}
	}

	max := 0
	for i := ii - 1; i >= 0; i-- {
		min := a[i][jj-1]
		for k := jj - 1; k >= 0; k-- {
			if max >= (i+1)*jj {
				return max
			}
			help := a[i][k]
			step := byte(0)
			t := true

			for j := k; j >= 0; j-- {
				if a[i][j] != 0 {
					step++
					if max < int(a[i][j]) {
						max = int(a[i][j])
					}
					if help > a[i][j] {
						help = a[i][j]
						if min > help {
							min = help
						}
					}
					if t && help != a[i][j] {
						k = j + 1
						t = false
					}
					if max < int(step)*int(help) {
						max = int(step) * int(help)
					}
				} else {
					t = false
					min = 0
					break
				}
			}
			if t {
				k = 0
			}
		}
		i = i - int(min)
	}

	return max
}

type List struct {
	value byte
	sum   int
	next  *List
}

func calculateRectangle(n byte, first *List) int {

	if first.next == nil {
		first.next = &List{n, int(n), nil}
		return int(n)
	}
	help := first
	max := 0
	put := false

	for help.next != nil && help.next.value <= n {
		help = help.next
		if help.value == n {
			put = true
		}
		help.sum += int(help.value)
		if max < help.sum {
			max = help.sum
		}
	}

	if !put && n != 0 {
		if help.next != nil {
			help = help.next
			help.sum = int(n) * (1 + help.sum/int(help.value))
			help.value = n
		} else {
			help.next = &List{n, int(n), nil}
			help = help.next
		}
		if max < help.sum {
			max = help.sum
		}
	}

	help.next = nil
	return max
}
func maximalRectangle(a [][]byte) int {
	var sum byte = 0
	jj, ii := len(a[0]), len(a)

	for j := 0; j < jj; j++ {
		sum = 0
		for i := 0; i < ii; i++ {
			sum = (a[i][j] - 48) * (sum + 1)
			a[i][j] = sum
		}
	}

	maxRectangle := 0
	for i := ii - 1; i >= 0; i-- {
		if maxRectangle >= (i+1)*jj {
			return maxRectangle
		}
		imin := a[i][jj-1]
		first := &List{0, 0, nil}

		for j := jj - 1; j >= 0; j-- {
			if imin > a[i][j] {
				imin = a[i][j]
			}
			if n := calculateRectangle(a[i][j], first); maxRectangle < n {
				maxRectangle = n
			}
		}

		i = i - int(imin)
	}

	return maxRectangle
}
func main() {
	fmt.Println(maximalRectangle([][]byte{
		{'1', '0', '1', '1', '0', '1'},
		{'1', '1', '1', '1', '1', '1'},
		{'0', '1', '1', '0', '1', '1'},
		{'1', '1', '1', '0', '1', '0'},
		{'0', '1', '1', '1', '1', '1'},
		{'1', '1', '0', '1', '1', '1'},
	}) == 8)
	fmt.Println(maximalRectangle([][]byte{
		{'1', '1', '0', '1'},
		{'1', '1', '0', '1'},
		{'1', '1', '1', '1'},
	}) == 6)
	fmt.Println(maximalRectangle([][]byte{
		{'0', '0', '1', '0'},
		{'1', '1', '1', '1'},
		{'1', '1', '1', '1'},
		{'1', '1', '1', '0'},
		{'1', '1', '0', '0'},
		{'1', '1', '1', '1'},
		{'1', '1', '1', '0'},
	}) == 12)
	fmt.Println(maximalRectangle([][]byte{
		{'0', '0', '1', '0'},
		{'0', '0', '1', '0'},
		{'0', '0', '1', '0'},
		{'0', '0', '1', '1'},
		{'0', '1', '1', '1'},
		{'0', '1', '1', '1'},
		{'1', '1', '1', '1'},
	}) == 9)

	fmt.Println(maximalRectangle([][]byte{
		{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
		{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
		{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
		{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
		{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
		{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
		{'1', '1', '1', '1', '1', '1', '1', '1', '1', '0', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
		{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
		{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
		{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
		{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
		{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
		{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
		{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
		{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
		{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
		{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
		{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
		{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
		{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
	}) == 260)

	fmt.Println(maximalRectangle([][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}) == 6)
	fmt.Println(maximalRectangle([][]byte{{'0'}}) == 0)

	fmt.Println(maximalRectangle([][]byte{{'1'}}) == 1)
}
