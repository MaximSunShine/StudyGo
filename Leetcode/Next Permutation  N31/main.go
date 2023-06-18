package main

import "fmt"

func nextPermutation(a []int) {
	find := false
	for i := len(a) - 1; i >= 0; i-- {
		if i == 0 || a[i-1] < a[i] {
			l, r := i, len(a)-1
			for ; l <= r; l, r = l+1, r-1 {
				if i != 0 && !find && a[i-1] < a[r] {
					a[i-1], a[r] = a[r], a[i-1]
					find = true
				}
				a[l], a[r] = a[r], a[l]
			}
			if l > r && !find && i != 0 {
				a[i-1], a[l] = a[l], a[i-1]
			}

			break
		}
	}
}
func main() {

	a := []int{3, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	nextPermutation(a)
	fmt.Println(a)

	a = []int{1, 2, 3}
	nextPermutation(a)
	fmt.Println(a)

	a = []int{3, 2, 1}
	nextPermutation(a)
	fmt.Println(a)

	a = []int{2, 3, 1}
	nextPermutation(a)
	fmt.Println(a)

	a = []int{4, 8, 8, 8, 9, 7, 7, 6, 4, 3}
	nextPermutation(a)
	fmt.Println(a)
}
