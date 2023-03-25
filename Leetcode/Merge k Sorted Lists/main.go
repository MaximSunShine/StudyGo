package main

import (
	"fmt"
)

type B struct {
	currentNum int
	length     int
}

func main() {
	var a [][]int = [][]int{{1, 2, 3, 4}, {2, 2, 3, 7}, {1, 1, 4, 4, 6, 8}}
	var b []B
	var c []int

	for _, v := range a {
		b = append(b, B{0, len(v)})
	}

	do := true
	for do {
		min := 0
		do = false
		for i := 0; i < len(b); i++ {
			if b[i].currentNum < b[i].length {
				min = a[i][b[i].currentNum]
				do = true
				break
			}
		}
		if do {
			for i, v := range b {
				if b[i].currentNum < b[i].length {
					if a[i][v.currentNum] < min {
						min = a[i][v.currentNum]
					}
				}
			}
			for i := 0; i < len(b); i++ {
				for b[i].currentNum < b[i].length {
					if a[i][b[i].currentNum] == min {
						c = append(c, min)
						b[i].currentNum++
					} else {
						break
					}
				}
			}
		}
	}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Print(c)
}
