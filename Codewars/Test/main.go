package main

import (
	"bytes"
	"fmt"
	"math"
	"unicode"
)

func SumOfDivided(lst []int) string {
	fmt.Println(lst)
	if len(lst) < 1 {
		return ""
	}

	maxNum := lst[0]
	for i := 0; i < len(lst); i++ {
		if maxNum < int(math.Abs(float64(lst[i]))) {
			maxNum = int(math.Abs(float64(lst[i])))
		}
	}

	//mAll := make(map[int]struct{})

	primeAll := []int{}
	if maxNum < 0 {
		maxNum *= -1
	}

	for i := 2; i <= maxNum; i++ {
		if isPrime(i) {
			primeAll = append(primeAll, i)
		}
	}

	mapMain := make(map[int]int)

	for i := 0; i < len(lst); i++ {
		for j := 0; j < len(primeAll) && primeAll[j] <= int(math.Abs(float64(lst[i]))); j++ {
			if lst[i]%primeAll[j] == 0 {
				mapMain[primeAll[j]] += lst[i]
			}
		}
	}

	s := ""
	for i := 0; i < len(primeAll); i++ {
		if v, ok := mapMain[primeAll[i]]; ok {
			s += fmt.Sprintf("(%v %v)", primeAll[i], v)
		}
	}
	fmt.Println("s = ", s)
	return s
}

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	//math.Abs()
	//m := make(map[int]struct{})
	if num == 2 {
		return true
	}

	if num%2 == 0 {
		return false
	}

	sqrtNum := int(math.Sqrt(float64(num)))
	for i := 3; i <= sqrtNum; i += 2 {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func Solution(a []int) (s string) {
	if len(a) == 0 {
		return ""
	}
	if len(a) == 1 {
		return fmt.Sprintf("%d", a[0])
	}

	//s := ""
	l, r := 0, 0

	for i := 1; i < len(a); i++ {
		if a[i]-a[i-1] == 1 {
			r = i
		} else {
			if r-l > 1 {
				s += fmt.Sprintf("%d-%d,", a[l], a[r])
			} else if r-l == 1 {
				s += fmt.Sprintf("%d,%d,", a[l], a[r])
			} else {
				s += fmt.Sprintf("%d,", a[l])
			}
			l, r = i, i
		}
	}
	if r-l > 1 {
		s += fmt.Sprintf("%d-%d,", a[l], a[r])
	} else if r-l == 1 {
		s += fmt.Sprintf("%d,%d,", a[l], a[r])
	} else {
		s += fmt.Sprintf("%d,", a[r])
	}

	return s[:len(s)-1]
}

func MultiplicationTable(size int) (rez [][]int) {
	// Implement me! :)
	for i := 1; i <= size; i++ {
		rez = append(rez, []int{})
		for j := 1; j <= size; j++ {
			rez[i-1] = append(rez[i-1], i*j)
		}
	}
	return
}

func CreatePhoneNumber(numbers [10]uint) string {
	b := []byte{}
	for i := 0; i < len(numbers); i++ {
		b = append(b, byte(numbers[i]))
	}
	return "(" + string(b[:3]) + ") " + string(b[3:6]) + "-" + string(b[6:])
}

func CamelCase(s string) string {
	//bytes.Replace()
	bb := bytes.Split([]byte(s), []byte(" "))
	for i, _ := range bb {
		bb[i][0] = bytes.ToUpper([]byte{bb[i][0]})[0]
	}

	return string(bytes.Join(bb, []byte{' '}))

	// your code here
}
func BreakCamelCase(s string) string {

	ss := ""
	first := 0

	if len(s) < 2 {
		return s
	}

	for i := 1; i < len(s); i++ {
		if unicode.IsUpper(rune(s[i])) {
			ss += s[first:i] + " "
			first = i
		}
	}

	if first < len(s)-1 {
		ss += s[first:]
	}

	return ss
}
func main() {
	BreakCamelCase("")
	BreakCamelCase("camel")
	BreakCamelCase("camelCasingTest")

	fmt.Println(SumOfDivided([]int{-4581, -2741, -4829, -809, -4999, -871, 4998, 1232, 3954, -4907, -3695, -4567, 1160, 3712, -1163}))
	for i := 0; i < 1000; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}

	//fmt.Println(Solution([]int{-10, -9, -8, -6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20}))
	//fmt.Println(CamelCase("say hello "))
}
