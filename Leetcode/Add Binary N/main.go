package main

import "fmt"

func plusOne(a string, pos int) (s string) {
	for i := pos; i >= 0; i-- {
		if a[i] == '1' {
			s = "0" + s
		} else {
			s = a[:i] + "1" + s
			return s
		}
	}

	return "1" + s
}
func addBinary(a string, b string) (s string) {
	if a[0] == '0' {
		return b
	}
	if b[0] == '0' {
		return a
	}

	if len(a) < len(b) {
		a, b = b, a
	}

	ar := []byte(a)

	var c byte = 0

	j := len(a) - len(b)

	for i := len(b) - 1; i >= 0; i-- {
		switch a[i+j] - 48 + b[i] - 48 + c {
		case 1:
			ar[i+j] = '1'
			c = 0
		case 2:
			ar[i+j] = '0'
			c = 1
		case 3:
			ar[i+j] = '1'
			c = 1
		}
	}

	if c == 1 {
		for i := j - 1; i >= 0; i-- {
			if a[i] == '1' {
				ar[i] = '0'
			} else {
				ar[i] = '1'
				return string(ar)
			}
		}
		return "1" + string(ar)
	}
	return string(ar)
}

func main() {
	fmt.Println(addBinary("0", "0"))
	fmt.Println(addBinary("0", "1"))
	fmt.Println(addBinary("1", "0"))
	fmt.Println(addBinary("1", "1"))
	fmt.Println(addBinary("11", "1"))
	fmt.Println(addBinary("1", "11"))
	fmt.Println(addBinary("1010", "1011"))
	fmt.Println(addBinary("1101010", "1011"))
	fmt.Println(addBinary("1111010", "1011"))
	fmt.Println(addBinary("100", "110010"))
	fmt.Println(addBinary("1111", "1111"))

}
