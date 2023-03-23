package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	s := ""
	find := false
	for i := 0; !find; i++ {
		if len(strs[0]) > i {
			current := string(strs[0][i])
			for _, a := range strs {
				if len(a) > i {
					if current != string(a[i]) {
						find = true
					}
				} else {
					find = true
				}
			}
			if !find {
				s += current
			}
		} else {
			find = true
		}
	}
	return s
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}
