package main

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {

	s := strings.Split(path, "/")
	rez := ""
	for i := 0; i < len(s); {

		switch s[i] {
		case "..":
			if i > 0 { // delete prev
				s = append(s[:i-1], s[i+1:]...)
				i = i - 1
			} else {
				s = append(s[:i], s[i+1:]...)
			}

		case ".", "":
			s = append(s[:i], s[i+1:]...)
		default:
			i++
		}
	}

	for _, val := range s {
		rez += "/" + val
	}

	if rez == "" {
		rez = "/"
	}

	return rez
}
func main() {
	fmt.Println(simplifyPath("/home/"))
	fmt.Println(simplifyPath("/../"))
	fmt.Println(simplifyPath("/home//foo/"))
	fmt.Println(simplifyPath(""))
	fmt.Println(simplifyPath("home/user/Documents/../../pro/.//...q/local/////bin"))

}
