package main

import (
	"bufio"
	"fmt"
	"os"
)

func reverse(s string) string {

	r := ""
	for i := 0; i < len(s); i++ {
		r = s[i:i+1] + r
	}

	return r
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	n, q := 0, 0

	s := []string{}
	help := ""
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &help)
		help = reverse(help)
		s = append(s, help)
	}

	t := []string{}
	fmt.Fscan(in, &q)
	for i := 0; i < q; i++ {
		fmt.Fscan(in, &help)
		help = reverse(help)
		t = append(t, help)
	}

	for _, v := range t {
		var test []string
		for _, vv := range s {
			if v[0] == vv[0] {
				if v != vv {
					test = append(test, vv)
				}
			}
		}

		if len(test) == 1 {
			fmt.Fprintln(out, reverse(test[0]))
			//fmt.Fprintln(out, reverse(v), " - ", reverse(test[0]))
			continue
		} else if len(test) == 0 {
			fmt.Fprintln(out, reverse(s[0]))
			//fmt.Fprintln(out, reverse(v), " - ", reverse(s[0]))
			continue
		} else {
			for i := 1; i < len(v); i++ {
				for j, vv := range test {
					if i >= len(vv) && len(test) > 1 {
						test = append(test[:j], test[j+1:]...)
					} else if vv[i] != v[i] && len(test) > 1 {
						test = append(test[:j], test[j+1:]...)
					}
					if len(test) == 1 {
						break
					}
				}
				if len(test) == 1 {
					fmt.Fprintln(out, reverse(test[0]))
					//fmt.Fprintln(out, reverse(v), " - ", reverse(test[0]))
					break
				}
			}
		}

	}

	//fmt.Fprintln(out, s, t)
}
