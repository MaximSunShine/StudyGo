package main

import "fmt"

func find(s string, a []string) (int, bool) {

	if len(s) > 0 {

		for i, val := range a {
			if s[:len(a[0])] == val {
				return i, true
				//is(s[3:], append(a[:i], a[i+1:]...))
			}
		}
	}
	return -1, false
}

func is(s string, a []string) bool {

	if len(a) == 0 {
		return true
	}
	ll := len(a[0])
	if len(s) < len(a)*ll && len(a) != 0 { // len(s)<ll
		return false
	}

	if i, ok := find(s[:ll], a); ok {
		return is(s[ll:], append(a[:i], a[i+1:]...))
	} else {
		return false
	}
}
func findSubstring(s string, words []string) (res []int) {
	n, m, k := len(s), len(words), len(words[0])
	cnt, window_cnt := make(map[string]int, m), 0
	for _, word := range words {
		cnt[word]++
	}
	for shift := 0; shift < k; shift++ {
		for i, j := shift, shift; j <= n; i, j = i+k, j+k {
			for ; j+k <= n && cnt[s[j:j+k]] > 0; j += k {
				cnt[s[j:j+k]]--
				window_cnt++
			}
			if window_cnt == m {
				res = append(res, i)
			}
			for ; i+k <= n && (j+k > n || s[i:i+k] != s[j:j+k]); i += k {
				cnt[s[i:i+k]]++
				window_cnt--
			}
		}
	}
	return
}
func findSubstring1(s string, c []string) []int {

	var b []int
	var a []string
	ll := len(c[0])
	for i := 0; i < len(s)-len(c)*ll+1; i++ {
		a = append([]string{}, c...)
		if is(s[i:], a) {
			b = append(b, i) //ababba
			help := s[i+1 : i+1+len(c)*len(c[0])]
			l := len(help)

			for i < len(s)-len(c)*ll && help[:l/2] == help[l/2:] {
				i++
				b = append(b, i)
			}
			for i < len(s)-len(c)*ll && s[i:i+ll] == s[i+ll*(len(c)-1)+1:i+ll*len(c)+1] {

			}
		}
	}

	return b
}
func main() {
	fmt.Println(findSubstring("abaababbaba", []string{"ba", "ab", "ab"}))
	fmt.Println(findSubstring("abcdef", []string{"abc", "def"}))
	fmt.Println(findSubstring("abcdefabcdef", []string{"a", "b", "c", "d", "e", "f"}))
	fmt.Println(findSubstring("aaaaaa", []string{"a", "a", "a"}))
	fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}))
	fmt.Println(findSubstring("barfoofoobarthefoobarman", []string{"bar", "foo", "the"}))
	fmt.Println(findSubstring("", []string{""}))
	fmt.Println(findSubstring("aaaaaa", []string{""}))
	fmt.Println(findSubstring("aaaaaa", []string{"a"}))
	fmt.Println(findSubstring("aaaaaa", []string{"aaaaaa"}))
	fmt.Println(findSubstring("aaaaaa", []string{"aaaaa"}))

}
