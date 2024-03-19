package main

import "fmt"

func makeKey(s string) string {
	var a = make([]byte, 26, 26)
	for _, v := range s {
		a[int(v-'a')]++
	}
	return string(a)
}

func groupAnagrams(strs []string) [][]string {
	set := make(map[string]int)
	var aa = make([][]string, 0, len(strs))
	for _, v := range strs {
		key := makeKey(v)
		if val, ok := set[key]; ok {
			aa[val] = append(aa[val], v)
		} else {
			set[key] = len(set)
			aa = append(aa, []string{v})
		}
	}

	return aa
}
func main() {
	fmt.Println(groupAnagrams([]string{"abc", "cba", "fgf", "gff"}))
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println(groupAnagrams([]string{""}))
	fmt.Println(groupAnagrams([]string{"a"}))
	fmt.Println(groupAnagrams([]string{"aaa"}))
}
