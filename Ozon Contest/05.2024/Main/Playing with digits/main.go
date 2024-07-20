package main

import (
	"math"
	"strconv"
)

func DigPow(n, p int) int {
	// your code
	s := strconv.Itoa(n)

	sum := 0
	for _, v := range s {
		rez := math.Pow(float64(v-48), float64(p))
		sum += int(rez)
		p++
	}

	if sum%n == 0 {
		return sum / n
	} else {
		return -1
	}
}
func main() {
	DigPow(89, 1)
}
