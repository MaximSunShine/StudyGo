package main

import (
	"fmt"
	"strconv"
)

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}

	second, first := 1, 1
	for i := 1; i < len(s); i++ {
		if s[i] == '0' {
			if s[i-1] > '2' || s[i-1] == '0' {
				return 0
			}
			second = 0
		}
		if s[i-1] > '2' || s[i-1] == '2' && s[i] > '6' {
			first = 0
		}
		first, second = second, first+second
	}
	return second
}
func numDecodings2(s string) int {
	if s[0] == '0' {
		return 0
	}

	zero := 0
	zero2 := 0
	for i := 1; i < len(s); i++ {
		if s[i] == '0' {
			zero++
			if s[i-1] == '1' || s[i-1] == '2' {
				zero2++
				s = s[:i-1] + s[i+1:]
				i--
			}
		}

	}
	if len(s) == 0 {
		return 1
	}

	first := 1
	if zero > 0 {
		first = 0
	}
	second := 1

	for i := 1; i < len(s); i++ {
		if s[i] == '0' {
			zero--
			if s[i-1] == '1' || s[i-1] == '2' {
				zero2--
			}
		}
		if s[i-1:i+1] == "00" {
			return 0
		}

		num, _ := strconv.Atoi(s[i-1 : i+1])
		if s[i-1] != '0' && num > 0 && num < 27 {
			if zero < 1 || zero == zero2 && i+1 < len(s) && s[i+1] != '0' {
				first, second = second, second+first
				//second--
			}
		} else {
			if first != 0 {
				first = second
			}
		}
	}
	if first == 0 {
		second = 0
	}

	fmt.Print(numDecodings1(s), "-")
	return second
}
func numDecodings1(s string) int {
	m := make(map[string]struct{})
	for i := 1; i < 27; i++ {
		m[strconv.Itoa(i)] = struct{}{}
	}

	ss := []string{}
	var f func(string)
	stop := false
	col := 0

	f = func(s string) {
		if stop {
			return
		}
		if len(s) == 0 {
			//fmt.Println(ss)
			col++
			return
		}
		for i := 1; i < 3; i++ {
			if len(s) >= i {
				if _, ok := m[s[:i]]; !ok {
					continue
				}
				ss = append(ss, s[:i])
				f(s[i:])
				ss = ss[:len(ss)-1]
			}
		}
	}

	f(s)

	return col
}
func numDecodings3(s string) int {
	//fmt.Print(numDecodings1(s), "-")
	if s[0] == '0' {
		return 0
	}

	first := 1
	second := 1
	for i := 1; i < len(s); i++ {
		if s[i] == '0' {
			if s[i-1] == '1' || s[i-1] == '2' {
				if i+1 < len(s) && s[i+1] == '0' {
					return 0
				}
				i++
				continue
			} else {
				return 0
			}
		}
		num, _ := strconv.Atoi(s[i-1 : i+1])
		if num > 0 && num < 27 {
			if i == len(s)-1 || i < len(s)-1 && s[i+1] != '0' {
				first, second = second, first+second
			}
		} else {
			first = second
		}
	}
	return second
}
func main() {

	fmt.Println(numDecodings("27"))
	fmt.Println(numDecodings("226"))
	fmt.Println(numDecodings("111111"))
	fmt.Println(numDecodings("100"))
	fmt.Println(numDecodings("10011"))
	fmt.Println(numDecodings("13"))
	fmt.Println(numDecodings("101110"))
	fmt.Println(numDecodings("25101"))
	fmt.Println(numDecodings("101010"))
	fmt.Println(numDecodings("10110"))
	fmt.Println(numDecodings("1010"))
	fmt.Println(numDecodings("2101"))
	fmt.Println(numDecodings("110"))
	fmt.Println(numDecodings("2611055971756562"))
	fmt.Println(numDecodings("230"))
	fmt.Println(numDecodings("207"))
	fmt.Println(numDecodings("1"))
	fmt.Println(numDecodings("131"))
	fmt.Println(numDecodings("1311"))
	fmt.Println(numDecodings("13111"))
	fmt.Println(numDecodings("131111"))
	//fmt.Println(numDecodings("1111111"))
	fmt.Println(numDecodings("31111111"))
	fmt.Println(numDecodings("13111111"))
	fmt.Println(numDecodings("13311111"))
	fmt.Println(numDecodings("13131111"))

	fmt.Println(numDecodings("12"))
	fmt.Println(numDecodings("226"))
	fmt.Println(numDecodings("06"))
	fmt.Println(numDecodings("0"))

}
