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
	t, col, help := 0, 0, 0
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		sum := 0
		fmt.Fscan(in, &col)
		m := make(map[int]int)
		for j := 0; j < col; j++ {
			fmt.Fscan(in, &help)
			m[help]++
		}

		for i, v := range m {
			sum += i * (v - v/3) // 5 - 1 = 4; 26 - 8 =  15
		}

		fmt.Fprintln(out, sum)
		//fmt.Fprintln(out, m)
	}

}
