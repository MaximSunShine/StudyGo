package main

import "fmt"

func addMinimum(word string) int {
	symbol := 'c'
	step := 0
	for _, v := range word {
		if symbol >= v {
			step++
		}
		symbol = v
	}

	return step*3 - len(word)
}

func addMinimum1(word string) int {

	step := 0
	if len(word) == 0 {
		word = "abc"
		return 3
	}
	if word[0] == 'b' {
		word = "a" + word
		step++
	} else if word[0] == 'c' {
		word = "ab" + word
		step += 2
	}

	for i := 0; step+i < len(word); {
		if word[step+i] == 'a' {
			if step+i == len(word)-1 {
				word += "bc"
				step += 2
				return step

			} else if word[step+i+1] != 'b' {
				word = word[:step+i+1] + "b" + word[step+i+1:]
				step++
			} else {
				i++
			}
		} else if word[step+i] == 'b' {
			if step+i == len(word)-1 {
				word += "c"
				step++
				return step

			} else if word[step+i+1] != 'c' {
				word = word[:step+i+1] + "c" + word[step+i+1:]
				step++
			} else {
				i++
			}
		} else if word[step+i] == 'c' {
			if step+i == len(word)-1 {
				return step
			} else if word[step+i+1] != 'a' {
				word = word[:step+i+1] + "a" + word[step+i+1:]
				step++
			} else {
				i++
			}
		} else {
			i++
		}
	}

	return 0
}
func main() {
	fmt.Println(addMinimum("b"))
	fmt.Println(addMinimum("aaa"))
	fmt.Println(addMinimum("abc"))
	fmt.Println(addMinimum("aaaabb"))
}
