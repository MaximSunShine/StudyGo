package main

import (
	"fmt"
)

func deleteDown(is *map[string]int, marki int, help string) {
	for i := marki; i >= 0; i-- {
		delete(*is, string(help[i]))

	}
}
func lengthOfLongestSubstring2(s string) int {
	res := 0
	if len(s) < 1 {
		return res
	}

	left := 0
	r := []rune(s)

	for right, current := range s {
		for i := left; i < right; i++ {
			if current == r[i] {
				left = i + 1
			}
		}

		if right-left+1 > res {
			res = right - left + 1
		}
	}

	return res
}
func lengthOfLongestSubstring(s string) int {
	longest := 0
	l := 0
	set := map[rune]int{}
	for r, ch := range s {
		if i, ok := set[ch]; ok {
			l = max(l, i+1)
		}
		set[ch] = r
		longest = max(longest, r-l+1)
	}
	return longest
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func lengthOfLongestSubstring1(s string) int {

	if s == "" {
		return 0
	}
	sum := 0
	min := 0
	is := make(map[string]int)
	is[string(s[0])] = 0
	length := 1
	var help string = s[:length]

	for i := 1; i < len(s); i++ {
		v, ok := is[string(s[i])]
		if !ok {
			length++
			is[string(s[i])] = i
			help = s[min : i+1]
		} else {
			if sum < length {
				sum = length
			}
			length = i - v
			deleteDown(&is, v-min, help) //delete(is, string(s[i]))
			is[string(s[i])] = i
			help = s[v+1 : i+1]
			min = v + 1

		}
	}
	if sum < length {
		sum = length
	}
	return sum
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcdefga"))
}
