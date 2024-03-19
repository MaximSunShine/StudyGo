package main

import (
	"fmt"
	"strings"
)

func fullJustify(s []string, maxWidth int) []string {
	rez := []string{}
	curLen := len(s[0])
	ch := make(chan []int)

	go func() {
		defer close(ch)
		if len(s) == 1 {
			ch <- []int{0, len(s[0])}
			return
		}
		for i := 1; i < len(s); i++ {
			if curLen == 0 {
				if i == len(s)-1 {
					ch <- []int{i, len(s[i])} //
					break
				}
				curLen = len(s[i])
				continue
			}

			if curLen+len(s[i])+1 <= maxWidth {
				curLen += len(s[i]) + 1
				if i == len(s)-1 {
					ch <- []int{i, curLen} //
					break
				}
				continue
			} else {
				i--
			}

			ch <- []int{i, curLen} //
			curLen = 0
		}
	}()

	l := 0
	for v := range ch {
		spaces := " "
		help := 0
		if v[0] != l {
			if v[0] != len(s)-1 {
				spaces = strings.Repeat(" ", (maxWidth-v[1])/(v[0]-l)+1)
				help = (maxWidth - v[1]) % (v[0] - l)
			}
		} else {
			rez = append(rez, s[l]+strings.Repeat(" ", maxWidth-v[1]))
			fmt.Println(rez[len(rez)-1], len(rez[len(rez)-1]))
			l++
			continue
		}
		rez = append(rez, s[l])
		for i := 1; i <= v[0]-l; i++ {
			if help > 0 {
				rez[len(rez)-1] += " "
				help--
			}
			rez[len(rez)-1] += spaces + s[l+i]

		}
		for len(rez[len(rez)-1]) < maxWidth {
			rez[len(rez)-1] += " "
		}

		fmt.Println(rez[len(rez)-1], len(rez[len(rez)-1]))
		l = v[0] + 1
	}

	return rez
}
func main() {
	fmt.Println(fullJustify([]string{"My", "momma", "always", "said,", "\"Life", "was", "like", "a", "box", "of", "chocolates.", "You", "never", "know", "what", "you're", "gonna", "get."}, 20))
	fmt.Println(fullJustify([]string{"a"}, 1))
	fmt.Println(fullJustify([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16))
	fmt.Println(fullJustify([]string{"What", "must", "be", "acknowledgment", "shall", "be"}, 16))
	fmt.Println(fullJustify([]string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}, 20))
}
