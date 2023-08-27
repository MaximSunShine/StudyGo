package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	sc.Scan()
	N, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	K, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	S := sc.Text()

	p := make([]int, N)
	for i := 0; i < N; i++ {
		sc.Scan()
		p[i], _ = strconv.Atoi(sc.Text())
	}

	d := make([]int, N)
	for i := 0; i < N; i++ {
		sc.Scan()
		d[i], _ = strconv.Atoi(sc.Text())
	}

	for i := 0; i < N; i++ {
		power := 1
		added := make(map[string]int)
		s := S[i : i+1]
		added[S[i:i+1]] = 1
		pos := i
		for j := 1; j < K; j++ {
			if _, is := added[S[p[pos]-1:p[pos]]]; is {
				added[S[p[pos]-1:p[pos]]]++
				l := ((int(S[p[pos]-1]) - 97) + (added[S[p[pos]-1:p[pos]]]-1)*d[p[pos]-1]) % 26
				s += string(97 + l)
			} else {
				added[S[p[pos]-1:p[pos]]] = 1
				power += 1
				s += S[p[pos]-1 : p[pos]]
			}
			pos = p[pos] - 1
		}
		fmt.Println(s)
	}

	fmt.Println(N, K, S, p, d)

}
