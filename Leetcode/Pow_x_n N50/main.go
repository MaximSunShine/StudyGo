package main

import (
	"fmt"
	"math"
)

func myPow(x float64, n int) float64 {

	/*if n < -1200 {
		return 0.0
	}*/

	if n == 0 {
		return 1.0
	}

	if x == 1 {
		return 1.0
	}

	if x == -1 {
		if n%2 == 0 {
			return 1
		}
		return -1
	}

	if x == 0 {
		if n == 0 {
			return 1
		}
		return 0
	}

	if n > 0 {
		return math.Pow(x, float64(n))
	}

	/*if n == -2147483648 {
		return 0
	}*/

	if n < 0 {
		rez := 1.0
		for ; rez != 0 && n < 0; n++ {

			rez /= x
		}
		return rez
	}

	return 0.0
}
func main() {

	fmt.Println(myPow(2.0, 10))

	fmt.Println(myPow(2.1, 3))

	fmt.Println(myPow(-2.0, -2))
	fmt.Println(myPow(-1.1, -2))

	fmt.Println(math.Pow(1.0000000000001, -2147483648))
} //0.99979
