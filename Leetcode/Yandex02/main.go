package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type S struct {
	index int
	val   int
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	sc.Scan()
	N, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	X, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	T, _ := strconv.Atoi(sc.Text())

	a := make([]S, 0, N)
	once := true
	for i := 0; i < N; i++ {
		sc.Scan()

		help, _ := strconv.Atoi(sc.Text())
		help -= X
		if help < 0 {
			help *= -1
		}

		if help-T <= 0 {

			if help-T < 0 {
				a = append(a, S{val: help, index: i + 1})

			} else if help-T == 0 && once {
				a = append(a, S{val: help, index: i + 1})
				once = false
			}
		}
	}

	t := true
	if len(a) > 1 {
		i := 1
		for t {
			t = false
			for j := i; j < len(a); j++ {
				if a[j-1].val > a[j].val {
					a[j], a[j-1] = a[j-1], a[j]
					t = true
				}
			}
		}
	}
	b := make([]S, 0, N)

	for i := 0; i < len(a) && T >= 0; i++ {
		T = T - a[i].val
		if T >= 0 {
			b = append(b, a[i])
		}
	}

	fmt.Println(len(b))
	for _, v := range b {
		fmt.Printf("%v ", v.index)
	}

}
