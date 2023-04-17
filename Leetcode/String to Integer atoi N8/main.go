package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

func myAtoi(s string) int {
	step := 0
	first := 0

	if len(s) == 0 {
		return 0
	}

	for s[step] == ' ' {
		step++
	}

	first = step

	if s[step] == '-' || s[step] == '+' {
		step++
	}

	for step < len(s) && s[step] != '.' && s[step] != ' ' && s[step] >= '0' && s[step] <= '9' {
		step++
	}

	s = s[first:step]

	help, err := strconv.ParseInt(s, 10, 32)

	if err != nil && len(s) > 0 {
		if numError, ok := err.(*strconv.NumError); ok {
			if numError.Err == strconv.ErrRange {
				if string(s[0]) == "-" {
					return math.MinInt32
				} else {
					return math.MaxInt32
				}
			}
		}
	}
	return int(help)
}

func myAtoi1(s string) int {
	reg, err := regexp.Compile("^ *([-+]?[0-9]+)")
	s = reg.FindString(s) //s = strings.TrimSpace(s)

	reg, err = regexp.Compile("[-+]?[0-9]+") //[ ]*
	s = reg.FindString(s)

	help, err := strconv.ParseInt(s, 10, 32)

	if err != nil {
		if numError, ok := err.(*strconv.NumError); ok {
			if numError.Err == strconv.ErrRange {
				if string(s[0]) == "-" {
					return math.MinInt32
				} else {
					return math.MaxInt32
				}
			}
		}
	}
	return int(help)
}
func main() {
	fmt.Println(myAtoi(""))
	fmt.Println(myAtoi(" a 21474836450a.58 nvajl "))
	fmt.Println(myAtoi("42"))
	fmt.Println(myAtoi("  -2147483648.1 nvajl"))
	fmt.Println(myAtoi("  -3.14159 nvajl"))
	fmt.Println(myAtoi("  +3.14159 nvajl"))

}
