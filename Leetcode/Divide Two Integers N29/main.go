package main

import "fmt"

func getSum(num int, target int, i int) (int, int) {

	if num+num < target {
		help, div := getSum(num+num, target, i+i)
		if help+num <= target {
			return help + num, div + i
		} else {
			return help, div
		}
	}

	return num, i
}

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}

	if divisor == 1 {
		return dividend
	}

	if divisor == -1 {
		help := 0 - dividend
		if help > 2147483647 {
			help = 2147483647
		}
		return help //
	}

	step := 0

	if divisor < 0 {
		divisor = 0 - divisor
		step++
	}
	if dividend < 0 {
		dividend = 0 - dividend
		step++
	}

	if divisor > dividend {
		return 0
	}

	if step == 2 {
		if dividend > 2147483647 {
			dividend = 2147483647
		}
	}

	v, i := getSum(divisor, dividend, 1)
	if dividend-v == divisor {
		i++
	}
	if step == 1 {
		i = 0 - i
	}

	return i
}
func main() {
	fmt.Println(divide(10, 3))           //3
	fmt.Println(divide(1004, -4))        //-251
	fmt.Println(divide(0, 1))            //0
	fmt.Println(divide(-1, 1))           //-1
	fmt.Println(divide(-4, 1))           //-4
	fmt.Println(divide(-1, 4))           //0
	fmt.Println(divide(1, -1))           //-1
	fmt.Println(divide(-4, -4))          //1
	fmt.Println(divide(-1000, -4))       //250
	fmt.Println(divide(-2147483648, -1)) //2147483647
	fmt.Println(divide(-2147483648, -4)) //536870911
	fmt.Println(divide(-2147483648, 2))  //-1073741824
	fmt.Println(divide(-2147483648, -2)) //1073741823
	fmt.Println(divide(-2147483648, 4))  //-536870912

}
