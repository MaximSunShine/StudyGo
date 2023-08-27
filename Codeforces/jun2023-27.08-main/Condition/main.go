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
	t, n := 0, 0

	fmt.Fscan(in, &t)

	var arr [][]int

	for i := 0; i < t; i++ {
		l, r := 15, 30

		fmt.Fscan(in, &n)
		arr = append(arr, make([]int, n, n))
		for idx, _ := range arr[i] {
			arr[i][idx] = -1

		}

		help := 0
		sign := ""
		find := false
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &sign, &help)

			if !find {
				if sign == ">=" {
					if l < help && help <= r {
						l = help
					}

					if help <= r {
						arr[i][j] = r
					} else {
						//arr[i][j] = -1
						find = true
					}

				} else { //sign == "<="
					if r > help && help >= l {
						r = help
					}
					if help >= l {
						arr[i][j] = r
					} else {
						//arr[i][j] = -1
						find = true
					}
				}
			}

		}

	}
	for _, v := range arr {
		for _, val := range v {
			fmt.Fprintln(out, val)
		}
		fmt.Fprintln(out)
	}
}
