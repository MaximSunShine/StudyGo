package main

import "fmt"

type list struct {
	str  int32
	pos  int
	next *list
}

func minWindow(s string, t string) string {

	m := make(map[int32]int)
	sMin := ""

	for _, v := range t {
		m[v]++
	}
	help := new(list)
	first := help
	step := len(t)

	for i, v := range s {
		if val, in := m[v]; in {
			m[v]--
			if val > 0 {
				step--
			}

			help.str = v
			help.pos = i
			help.next = new(list)
			help = help.next

			for step == 0 {
				for m[first.str] < 0 {
					m[first.str]++
					first = first.next
				}

				if len(s[first.pos:i+1]) == len(t) {
					return s[first.pos : i+1]
				}
				if sMin == "" || len(sMin) > len(s[first.pos:i+1]) {
					sMin = s[first.pos : i+1]
				}
				m[first.str]++
				step++
				first = first.next
			}
		}
	}
	return sMin
}
func main() {
	fmt.Println(minWindow("aaabdabcefaecbef", "abc"))
	fmt.Println(minWindow("ab", "b"))
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
	fmt.Println(minWindow("a", "a"))
	fmt.Println(minWindow("a", "aa"))
}
