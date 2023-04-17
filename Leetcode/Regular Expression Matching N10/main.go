package main

import (
	"fmt"
	"regexp"
)

func isMatch(s string, p string) bool {
	/*
		reg, err := regexp.Compile(p)
		if err != nil {
			return false
		}
		return s == reg.FindString(s)*/
	reg, _ := regexp.MatchString(p, s)
	return reg
}

func main() {

	fmt.Println("false -> ", isMatch("aa", "a"))
	fmt.Println("true -> ", isMatch("aa", "a*"))
	fmt.Println("true -> ", isMatch("ab", ".*"))

}
