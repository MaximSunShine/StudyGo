package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	n, t := 0, 0
	fmt.Fscan(in, &n, &t)

	m := make(map[byte]int)
	for i := 0; i < n; i++ {
		help := ""
		fmt.Fscan(in, &help)
		m[help[0]]++
	}

Loop:
	for i := 0; i < t; i++ {
		mTest := make(map[byte]int)
		s := ""
		fmt.Fscan(in, &s)
		if n == len(s) {
			for j := 0; j < len(s); j++ {
				mTest[s[j]]++
			}
		} else {
			fmt.Fprintln(out, "NO")
			continue
		}

		if len(m) != len(mTest) {
			fmt.Fprintln(out, "NO")
			continue
		}

		for key, value := range m {
			if mTest[key] == value {
				continue
			} else {
				fmt.Fprintln(out, "NO")
				continue Loop
			}
		}
		fmt.Fprintln(out, "YES")
	}
}
