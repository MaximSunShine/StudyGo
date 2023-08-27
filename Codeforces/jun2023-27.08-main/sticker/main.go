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

	s := ""
	fmt.Fscan(in, &s)
	n := 0
	fmt.Fscan(in, &n)

	f, l := 0, 0
	t := ""

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &f, &l, &t)
		s = s[:f-1] + t + s[l:]
	}

	fmt.Fprintln(out, s)
}
