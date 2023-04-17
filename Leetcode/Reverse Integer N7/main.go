package main

import (
	"math"
	"strconv"
)

const (
	minInt32 = int64(math.MinInt32) // -1 << 31
	maxInt32 = int64(math.MaxInt32) // 1 << 31 -1
)

func reverse(n int) int {

	a := strconv.Itoa(n)
	minus := ""
	b := ""
	if string(a[0]) == "-" {
		minus = string(a[0])
		a = a[1:]
	}

	for _, val := range a {
		b = string(val) + b
	}

	b = minus + b

	help, err := strconv.ParseInt(b, 10, 32)

	if err != nil {
		help = 0
	}

	return int(help)
}

func main() {
	//2147483647с
	//1563847412
	//-2147483648
	//-1563847412
	var n = 1563847412
	n = reverse(n)
	println(n)
}
