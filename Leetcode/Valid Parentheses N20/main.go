package main

import "fmt"

// 40 41 ()
// 91 93 []
// 123 125 {}
// (] = 133  (} = 165 [) = 132 [} = 216 {) = 164 {] = 216

func isValid1(s string) bool {
	if len(s)%2 == 1 {
		return false
	}

	pairs := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	stack := []rune{}

	for _, r := range s {
		if _, ok := pairs[r]; ok {
			stack = append(stack, r)
		} else if len(stack) == 0 || pairs[stack[len(stack)-1]] != r {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func isValid(s string) bool {

	if len(s)%2 != 0 {
		return false
	}

	i := 0

	for i >= 0 && i < len(s)-1 {
		help := s[i+1] + s[i]
		if (help == 81 || help == 184 || help == 248) && s[i] < s[i+1] {
			s = s[:i] + s[i+2:]
			if i > 0 {
				i = i - 1
			}
			continue
		}
		if help == 133 || help == 165 || help == 132 || help == 216 || help == 164 {
			return false
		}

		i++
	}

	if s == "" {
		return true
	}
	return false
}
func main() {
	fmt.Println(isValid("[]{"))
	fmt.Println(isValid("()"))
	fmt.Println(isValid("[](){}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("[({})]"))
	fmt.Println(isValid("[({}])"))
	fmt.Println(isValid("["))
	fmt.Println(isValid1("[({}])"))
	fmt.Println(isValid1("[({})]"))

}
