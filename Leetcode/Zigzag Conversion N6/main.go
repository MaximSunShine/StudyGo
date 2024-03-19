package main

import "fmt"

func convert(s string, numRows int) string {

	if len(s) <= numRows || numRows == 1 {
		return s
	}
	ns := ""
	if 2*numRows > len(s)+1 {
		z := 2*numRows - (len(s) + 1)
		numRows -= z
		ns = s[:z]
		s = s[z:]
	}

	help := make([][]uint8, numRows, numRows)
	step := 0
	for step < len(s) {

		for k := 0; k < numRows && step < len(s); k++ {
			help[k] = append(help[k], s[step])
			step++
		}
		for k := numRows - 2; k > 0 && step < len(s); k-- {
			help[k] = append(help[k], s[step])
			step++
		}

	}

	for i := 0; i < numRows; i++ {
		ns += string(help[i])
	}
	return ns
}

func convert1(s string, numRows int) string {

	if len(s) <= numRows || numRows == 1 {
		return s
	}

	help := make([]string, numRows, numRows)

	for step := 0; step < len(s); {

		for k := 0; k < numRows && step < len(s); k++ {
			help[k] += string(s[step])
			step++
		}
		for k := numRows - 2; k > 0 && step < len(s); k-- {
			help[k] += string(s[step])
			step++
		}

	}

	ns := ""
	for i := 0; i < numRows; i++ {
		ns += help[i]
	}
	return ns
}

func main() {
	fmt.Println(convert("f", 3))
	fmt.Println(convert("PAYPALISHIRING", 1))
	fmt.Println(convert("PAYPALISHIRING", 2))
	fmt.Println(convert("PAYPALISHIRING", 3))
	fmt.Println(convert1("PAYPALISHIRING", 4))
	fmt.Println(convert("PAYPALISHIRING", 4))

	fmt.Println(convert1("1234567890abcdefghi", 13))
	fmt.Println(convert("1234567890abcdefghi", 13))

	fmt.Println(convert1("1234567890abcdefghim", 14))
	fmt.Println(convert("1234567890abcdefghim", 14))
}
