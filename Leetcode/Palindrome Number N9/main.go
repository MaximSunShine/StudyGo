package main

import (
	"fmt"
)

func isPalindrome(x int) bool {

	if x < 0 { // убираем все отрицательные
		return false
	}
	if x >= 0 && x < 10 { // убираем одноразрядные
		return true
	}

	first := 10
	last := 10

	for x/last > 9 {
		last = last * 10
	}

	for first <= last {
		if x/last != x%first { //1234 -> 1 != 4 если хвосты одинаковые, укорачиваем, иначе ложь
			return false
		}
		x = (x - last*(x/last)) / 10 //укорачивание
		last = last / 100
	}
	return true
}
func main() {
	fmt.Println(isPalindrome(10))
	fmt.Println(isPalindrome(100))
	fmt.Println(isPalindrome(1000))
	fmt.Println(isPalindrome(8))
	fmt.Println(isPalindrome(0))
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(1221))
	fmt.Println(isPalindrome(1235321))
	fmt.Println(isPalindrome(12354321))
}
