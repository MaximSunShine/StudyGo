package main

import "fmt"

var romanencoder map[string]int

func initRomanencoder() {
	romanencoder = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
}

func romanToInt(s string) int {

	current := romanencoder[string(s[len(s)-1])]
	sum := current

	for i := len(s) - 2; i >= 0; i-- {
		next := romanencoder[string(s[i])]

		if current > next {
			sum -= next
		} else {
			sum += next
			current = next
		}
	}

	return sum
}
func main() {
	initRomanencoder()

	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))

}
