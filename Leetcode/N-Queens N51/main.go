package main

import (
	"fmt"
	"sync"
)

var step int

func Print(a []string) {
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
	fmt.Println()
}

func fillDots(n int) []string {
	var a = make([]string, n, n)
	s := ""
	for i := 0; i < n; i++ {
		s += "."
	}

	for i := 0; i < n; i++ {
		a[i] = s
	}
	return a
}

func fillBroken(a []string, i int, bottom int, val byte) {
	for j := 0; j+bottom < len(a); j++ {
		b := []byte(a[j+bottom])

		b[i] = val
		if i-(j+1) >= 0 {
			b[i-(j+1)] = val
		}
		if i+(j+1) < len(a[0]) {
			b[i+(j+1)] = val
		}
		a[j+bottom] = string(b)

	}
}

func findSolution(n int, wg *sync.WaitGroup, ch chan []string) {
	defer wg.Done()
	var a = fillDots(n)
	//Print(a)

	var putQueen func([]string, int)
	putQueen = func(a []string, bottom int) {

		for i := 0; i < len(a[0]); i++ {
			if a[bottom][i] == '.' {
				a[bottom] = a[bottom][:i] + "Q" + a[bottom][i+1:]
				var aa = make([]string, len(a), len(a))
				copy(aa, a)
				fillBroken(aa, i, bottom+1, '*') //Broke
				//Print(a)
				if len(a) == bottom+1 {
					step++
					ch <- aa
					//Print(aa)
					return
				}
				putQueen(aa, bottom+1)
				a[bottom] = a[bottom][:i] + "." + a[bottom][i+1:]
				fillBroken(aa, i, bottom, '.') //Unbroke
			}
		}
	}
	putQueen(a, 0)

}

func solveNQueens(n int) [][]string {
	if n == 1 {
		return [][]string{{"Q"}}
	}

	if n < 4 {
		return [][]string{}
	}

	var ch = make(chan []string)
	var wg sync.WaitGroup
	wg.Add(1)
	go findSolution(n, &wg, ch)

	go func() {
		wg.Wait()
		close(ch)
	}()

	var rez [][]string
	for v := range ch {
		for i := 0; i < len(v); i++ {
			s := ""
			for j := 0; j < len(v[0]); j++ {
				if v[i][j] == '*' {
					s += "."
				} else {
					s += string(v[i][j])
				}
			}
			v[i] = s
		}
		rez = append(rez, v)
	}

	fmt.Println(step)
	return rez
}

func main() {
	for i := 0; i < 10; i++ {
		step = 0
		solveNQueens(i)
	}
	//fmt.Println(step)
	fmt.Println(solveNQueens(6))

}
