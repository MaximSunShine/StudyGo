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
	n := 0
	fmt.Fscan(in, &n)
LOOP:
	for i := 0; i < n; i++ {
		s := ""
		fmt.Fscan(in, &s)
		if len(s) == 1 {
			fmt.Fprintln(out, "YES")
			continue
		}
		if s[0] != s[len(s)-1] {
			fmt.Fprintln(out, "NO")
			continue
		}

		for j := 1; j < len(s); j++ {
			if s[j] == s[0] {
				continue
			}
			if s[j] != s[j-1] {
				if s[j+1] != s[0] {
					fmt.Fprintln(out, "NO")
					continue LOOP
				}
			}
		}
		fmt.Fprintln(out, "YES")
	}
}
