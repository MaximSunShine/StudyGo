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

	n, q := 0, 0
	fmt.Fscan(in, &n, &q)

	m := make(map[int]int)
	text, id, numSms := 0, 0, 1
	for i := 0; i < q; i++ {
		//fmt.Fscanln(in,[]int{})
		fmt.Fscan(in, &text, &id)
		if text == 1 {
			if id == 0 {
				for j := 1; j < n; j++ {
					m[j] = numSms
				}
			} else {
				m[id] = numSms
			}
			numSms++
		} else { //text == 2
			if value, inMap := m[id]; inMap {
				fmt.Fprintln(out, value)
			} else {
				fmt.Fprintln(out, 0)
			}
		}
	}
}
