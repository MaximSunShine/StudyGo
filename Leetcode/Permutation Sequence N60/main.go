package main

import (
	"fmt"
)

var aa = []rune{'1', '2', '3', '4', '5', '6', '7', '8', '9'}
var b []int

func initial() {
	if b == nil {
		b = make([]int, 9, 9)
		sum := 1
		for i := 1; i <= len(b); i++ {
			sum *= i
			b[i-1] = sum
		}
	}
}

func getPermutation(n int, step int) string {
	a := make([]rune, 9, 9)
	copy(a, aa)

	initial()

	a = a[:n]
	if step == b[n-1] { //if step is the last combination
		for i := 0; i < n/2; i++ { //reverse array
			a[i], a[n-i-1] = a[n-i-1], a[i]
		}
		return string(a)
	}

	s := ""
	for i := n - 1; i > 0; i-- {
		if step == 1 { //have not more combination
			return s + string(a[:len(a)])
		} else if step < b[i-1] { //go next without calc
			s += string(a[0])
			a = append(a[:0], a[1:]...)
		} else {
			pos := (step - 1) / b[i-1]
			step = step - pos*b[i-1] // or help := (step-1)%b[i-1] + 1
			s += string(a[pos])
			a = append(a[:pos], a[pos+1:]...)
		}
	}
	return s + string(a[0])
}
func main() {

	fmt.Println(getPermutation(4, 2))
	fmt.Println(getPermutation(3, 2))
	fmt.Println(getPermutation(3, 3))
	fmt.Println(getPermutation(4, 9))
	fmt.Println(getPermutation(3, 1))
	fmt.Println(getPermutation(4, 24))
	fmt.Println(getPermutation(3, 5))
	fmt.Println(getPermutation(9, 23))

	fmt.Println(getPermutation(4, 12))
	fmt.Println(getPermutation(5, 5*24/2))

}
