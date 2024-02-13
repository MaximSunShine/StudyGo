package main

import "fmt"

func isMatch1(s string, p string) bool {

	next := 0
	i := 0
	for ; i < len(p); i++ {
		if p[i] == '?' || p[i] == s[next] {
			next++
			continue
		}
		if p[i] == '*' {
			for i < len(p) && p[i] == '*' { // move -> ******abc = abc
				if i == len(p)-1 {
					return true
				}
				i++
			}
			for next < len(s)-1 && p[i] != s[next] && p[i] != '?' {
				next++
			}

			if i < len(p)-1 && next < len(s)-1 && isMatch(s[next+1:], "*"+p[i:]) == true {
				return true
			}

		}
		if next < len(s) && i < len(p) && p[i] != s[next] {
			break
		}
		if next < len(s) {
			next++
		} else {
			break
		}
	}
	i--
	if i < len(p)-1 || next < len(s)-1 {
		return false
	}

	return true
}
func isMatch2(s string, p string) bool {

	for ; len(p) > 0; p = p[1:] {
		if p[0] == '*' {
			for len(p) > 0 && p[0] == '*' { // move -> ******abc = abc
				if len(p) == 1 {
					return true
				}
				p = p[1:]
			}

			k := 0
			for i, v := range p {
				if v == '*' || v == '?' {
					k = i
					break
				}
			}

			for len(s) > k && p[:k] != s[:k] {
				s = s[1:]
			}

			if len(s) == 0 {
				break
			}

			if len(p) > 0 && len(s) > 0 && isMatch(s[1:], "*"+p) == true {
				return true
			}

		}
		if len(s) == 0 {
			break
		}
		if p[0] == '?' || p[0] == s[0] {
			s = s[1:]
			continue
		}
		if len(s) > 0 && len(p) > 0 && p[0] != s[0] {
			break
		}
		if len(s) > 0 {
			s = s[1:]
		} else {
			break
		}
	}

	if len(p) > 0 || len(s) > 0 {
		return false
	}

	return true
}

func Q(s, p string) bool {

	if len(p) == 0 {
		return true
	}

	k := 0
	for k = 0; k < len(p) && p[k] != '?'; k++ { //aaa*aa or aaa?aa => aaa
	}

	if k == 0 && len(s) > 0 {
		return true
	}

	if p[:k] != s[:k] {
		return false
	} else {
		if len(p) > k && p[k] == '?' {
			k++
		}
		return Q(s[k:], p[k:])
	}
}
func isMatch(s string, p string) bool {
	k := 0
	for len(p) > 0 {
		k = 0
		if p[0] == '*' {
			for ; k < len(p)-1 && p[k] == '*'; k++ {
			}
			p = p[k:]
			if len(p) == 1 && p[0] == '*' {
				return true
			}

			if len(p) == 1 && p[0] == '?' && len(s) > 0 {
				return true
			}

			if p[0] == '?' {
				if len(s) > 0 {
					s = s[1:]
					if len(p) > 1 {
						p = p[1:]
					}
				} else {
					return false
				}
			}

			f := false
			for k = 0; k < len(p) && p[k] != '*'; k++ { //aaa*aa or aaa?aa => aaa
				if p[k] == '?' {
					f = true
				}
			}

			if f {
				x := 0
				for x = 0; x < len(p) && p[x] != '?'; x++ {
				}
				for len(s) > k {

					if len(p) == 1 && p[0] == '*' {
						return true
					}

					for ; k < len(s) && p[:x] != s[:x]; s = s[1:] {
					}

					if p[:x] == s[:x] && p[0] != '?' && Q(s[x+1:], p[x+1:k]) {
						p = p[k:]
						s = s[k:]
						break
					} else {
						if len(s) > len(p) {
							s = s[1:]
							if len(p) > 1 {
								p = p[1:]
							}
						} else {
							return false
						}
					}
				}
				continue
			}

			for ; k < len(s) && p[:k] != s[:k]; s = s[1:] {
			}

			if k == len(p) && p[len(p)-1] != '*' {
				if len(s) >= k && Q(s[len(s)-len(p):], p) {
					return true
				}
				return false
			} else {
				if len(s) < k || p[:k] != s[:k] {
					return false
				} else {
					p = p[k:]
					s = s[k:]
				}
				continue
			}
		}

		if len(s) == 0 {
			return false
		}

		if len(s) > 0 && (p[0] == '?' || p[0] == s[0]) {
			s = s[1:]
			p = p[1:]
			continue
		}
		return false

	}

	if len(s) > 0 || len(p) > 0 {
		return false
	}

	return true
}

func main() {

	fmt.Println(isMatch("bbbab", "*??a?") && true) //true

	fmt.Println(!(isMatch("bababb", "**??*") && false)) //false

	fmt.Println(!(isMatch("bbbaba", "bb**??") && false)) //false

	fmt.Println(!(isMatch("aabab", "?**aa?") && false)) //false

	fmt.Println(!(isMatch("abcdefghij", "abc*defghijk") && false)) //false

	fmt.Println(!(isMatch("b", "?*?") && false)) //false

	fmt.Println(isMatch("hi", "*?") && true) //true

	fmt.Println(!(isMatch("mississippi", "m??*ss*?i*pi") && false)) //false

	fmt.Println(isMatch("adceb", "*a*b") && true) //true

	fmt.Println(isMatch("abbabaaabbabbaababbabbbbbabbbabbbabaaaaababababbbabababaabbababaabbbbbbaaaabababbbaabbbbaabbbbababababbaabbaababaabbbababababbbbaaabbbbbabaaaabbababbbbaababaabbababbbbbababbbabaaaaaaaabbbbbaabaaababaaaabb",
		"*?aa*ba?*a*bb*aa*ab*a*aaa?aaa*a*aaaa*bbabb*b*b*aaaaaaaaa*a*ba*bbb*a*ba*bb*bb**a*b*bb?"), false)

	fmt.Println(isMatch("", "*****"), true)
	fmt.Println(isMatch("acbcbcbcbcb", "a*c?*cb") && true) //true
	fmt.Println(isMatch("acbcbcbcbcb", "a*c?*cb") && true) //true
	fmt.Println(isMatch("ababa", "a*a?a") && true)         //true
	fmt.Println(isMatch("abcdaba", "a*a?a") && true)       //true
	fmt.Println(isMatch("aacccaacda", "aa*aa??a") && true) //true

	fmt.Println((isMatch("aacccaacda", "aa****aa??a") && true)) //true

	fmt.Println(!(isMatch("aa", "a") && false))  //false
	fmt.Println(isMatch("aa", "*") && true)      //true
	fmt.Println(!(isMatch("cb", "?a") && false)) //false

	fmt.Println(isMatch("acacacacacb", "a*cb") && true) //true
}
