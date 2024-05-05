package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var in *bufio.Reader
	in = bufio.NewReader(os.Stdin)
	var out *bufio.Writer
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()
	n, m := 0, 0

	fmt.Fscan(in, &n, &m)

	if n == 1 {
		help := 0
		fmt.Fscan(in, &help)
		if help < m {
			fmt.Fprint(out, help+1)
			return
		} else {
			fmt.Fprint(out, -1)
			return
		}
	}

	if m == 1 || n == m {
		fmt.Fprintln(out, -1)
		return
	}

	mm := make([]int, n, n)

	for i := n - 1; i >= 0; i-- {
		mm[i] = m - n + i
	}
	rez := make([]int, 0, n)
loop:
	for i := 0; i < n; i++ {
		help := 0
		fmt.Fscan(in, &help)
		if help == len(mm) {
			fmt.Fprintln(out, -1)
			return
		}

		for j := help; j < m; j++ {
			if j == mm[j] {
				rez = append(rez, j+1)
				mm[help] = j + 1
				mm[j] = j + 1
				continue loop
			} else if j < mm[j] {
				if mm[j] == m {
					fmt.Fprintln(out, -1)
					return
				}
				j = mm[j] - 1
				continue
			}
		}
		fmt.Fprintln(out, -1)
		return
	}
	for _, v := range rez {
		fmt.Fprint(out, v, " ")
	}
}
