package main

import (
	"fmt"
	"strconv"
)

func grayCode(n int) []int {
	code_count := 1 << n
	rez := make([]int, code_count)

	for i := 0; i < code_count; i += 1 {
		mask := i >> 1
		rez[i] = i ^ mask
	}

	return rez
}

func doCode(a []int) {
	for i := 0; i < len(a); i++ {
		if a[i] == 0 {
			a[i] = 1
		} else {
			a[i] = 0
		}

		doCode(a)
	}
}

func grayCode1(n int) []int {

	step := 1 << n
	a := make([]byte, n, n)
	m := make(map[string]struct{})

	rez := []int{0}

	for i := 0; i < len(a); i++ {
		a[i] = '0'
	}
	word := string(a)

	var f func([]byte) bool
	f = func(a []byte) bool {
		for i := 0; i < len(a); i++ {
			if a[i] == 48 {
				a[i]++
			} else {
				a[i]--
			}

			if len(m) == step && string(a) == word {
				//fmt.Println("Done ->", rez)
				return true
			}

			if _, is := m[string(a)]; !is {
				m[string(a)] = struct{}{}
				help, _ := strconv.ParseInt(string(a), 2, 32)
				rez = append(rez, int(help))
				if f(a) {
					return true
				}
				delete(m, string(a))
				rez = rez[:len(rez)-1]
			}
			if a[i] == 48 {
				a[i]++
			} else {
				a[i]--
			}
		}
		return false
	}
	m[string(a)] = struct{}{}
	f(a)

	return rez
}
func main() {
	fmt.Println(grayCode(2))

}
