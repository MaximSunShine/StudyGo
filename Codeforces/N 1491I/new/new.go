package main

import "fmt"

type animal struct {
	num   int
	power [3]int
}

func doKing(a []animal) (int, int) {
	col, count, pow, i, j := 0, 0, 0, 0, 1

	for count < len(a) || col < 100000 {
		if i == len(a)-1 {
			j = 0
		}
		//fmt.Println(a[i].power[pow], a[j].power[0])
		if a[i].power[pow] > a[j].power[0] {
			count = 0
			pow++
			if pow == 3 {
				return a[i].num, col + 1
			}
			a[i].num, a[i].power, a[j].num, a[j].power = a[j].num, a[j].power, a[i].num, a[i].power
		} else {
			pow = 1
		}
		if i == len(a)-1 {
			i = 0
		} else {
			i++
		}
		j++
		col++
		count++
	}
	return -1, -1
}

func main() {
	var n int
	fmt.Scan(&n)

	var a = make([]animal, n, n)

	for i := 0; i < len(a); i++ {
		a[i].num = i
		fmt.Scan(&a[i].power[0], &a[i].power[1], &a[i].power[2])
	}

	//fmt.Println(a)
	king, colMove := doKing(a)
	fmt.Printf("%v %v", king, colMove)

	//fmt.Println(a)

}
