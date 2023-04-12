package main

import (
	"fmt"
)



func fizzBuzz(n int) []string {
	arr := make([]string, n)

	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			if i%5 == 0 {
				arr[i-1] = "FizzBuzz"
				continue
			}
			arr[i-1] = "Fizz"
			continue
		}

		if i%5 == 0 {
			arr[i-1] += "Buzz"
			continue
		}

		if arr[i-1] == "" {
			arr[i-1] = fmt.Sprintf("%d", i)
		}
	}
	return arr
}

func main() {

	fmt.Println(fizzBuzz(15))
}
