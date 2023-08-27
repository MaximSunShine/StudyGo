package main

import "fmt"

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func longestValidParentheses(s string) int {
	stack := []int{-1}
	maxLength := 0

	for i, char := range s {
		if char == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				maxLength = Max(maxLength, i-stack[len(stack)-1])
			}
		}
	}

	return maxLength
}
func rec(s string, i, bottom int) (int, int, int) {

	mainSum := 0
	sum := 0
	for ; i < len(s); i++ {
		if s[i] == 40 {
			//bottom++
			if i+1 < len(s) {
				help := 0
				bottomHelp := 0
				i, bottomHelp, help = rec(s, i+1, bottom+1)
				if bottomHelp != bottom {
					if sum < help {
						sum = help
					}
				} else {
					sum += help
				}
			}
		} else if bottom != 0 {
			bottom--
			sum++
			return i, bottom, sum
		} else {
			if mainSum < sum {
				mainSum = sum
			}
			sum = 0
		}
	}

	if mainSum < sum {
		mainSum = sum
	}

	return i, bottom, mainSum
}

func longestValidParentheses1(s string) int {
	_, _, i := rec(s, 0, 0)
	return i * 2
}
func longestValidParentheses2(s string) int {
	// ( - 40
	// ) - 41
	sum := 0
	step := 0
	output := 0
	help := 0
	for _, v := range s {
		if v == 40 {
			sum++
		} else {
			sum--
			if sum == 0 {
				if help < step+1 {
					help = step + 1
				}
			}
			if sum < 0 {
				if output < step {
					output = step
				}
				step = 0
				sum = 0
				continue
			}
		}

		if sum >= 0 {
			step++
		}

	}
	step = step - sum
	if sum > 0 {
		if help > step-help {
			return help
		} else {
			return step - help
		}
	}

	if output < step {
		output = step
	}

	return output
}
func main() {

	fmt.Println(longestValidParentheses(")("), " - 0")
	fmt.Println(longestValidParentheses("()"), " - 2")

	fmt.Println(longestValidParentheses("((("), " - 0")
	fmt.Println(longestValidParentheses("(()"), " - 2")
	fmt.Println(longestValidParentheses("()("), " - 2")
	fmt.Println(longestValidParentheses("())"), " - 2")

	fmt.Println(longestValidParentheses(")(("), " - 0")
	fmt.Println(longestValidParentheses(")()"), " - 2")
	fmt.Println(longestValidParentheses("))("), " - 0")
	fmt.Println(longestValidParentheses(")))"), " - 0")

	fmt.Println(longestValidParentheses("(((("), " - 0")
	fmt.Println(longestValidParentheses("((()"), " - 2")
	fmt.Println(longestValidParentheses("(()("), " - 2")
	fmt.Println(longestValidParentheses("(())"), " - 4")
	fmt.Println(longestValidParentheses("()(("), " - 2")
	fmt.Println(longestValidParentheses("()()"), " - 4")
	fmt.Println(longestValidParentheses("())("), " - 2")
	fmt.Println(longestValidParentheses("()))"), " - 2")

	fmt.Println(longestValidParentheses(")((("), " - 0")
	fmt.Println(longestValidParentheses(")(()"), " - 2")
	fmt.Println(longestValidParentheses(")()("), " - 2")
	fmt.Println(longestValidParentheses(")())"), " - 2")
	fmt.Println(longestValidParentheses("))(("), " - 0")
	fmt.Println(longestValidParentheses("))()"), " - 2")
	fmt.Println(longestValidParentheses(")))("), " - 0")
	fmt.Println(longestValidParentheses("))))"), " - 0")

	fmt.Println(longestValidParentheses("((((("), " - 0")
	fmt.Println(longestValidParentheses("(((()"), " - 2")
	fmt.Println(longestValidParentheses("((()("), " - 2")
	fmt.Println(longestValidParentheses("((())"), " - 4")
	fmt.Println(longestValidParentheses("(()(("), " - 2")
	fmt.Println(longestValidParentheses("(()()"), " - 4")
	fmt.Println(longestValidParentheses("(())("), " - 4")
	fmt.Println(longestValidParentheses("(()))"), " - 4")
	fmt.Println(longestValidParentheses("()((("), " - 2")
	fmt.Println(longestValidParentheses("()(()"), " - 2!")
	fmt.Println(longestValidParentheses("()()("), " - 4")
	fmt.Println(longestValidParentheses("()())"), " - 4")
	fmt.Println(longestValidParentheses("())(("), " - 2")
	fmt.Println(longestValidParentheses("())()"), " - 2!")
	fmt.Println(longestValidParentheses("()))("), " - 2")
	fmt.Println(longestValidParentheses("())))"), " - 2")

	fmt.Println(longestValidParentheses(")(((("), " - 0")
	fmt.Println(longestValidParentheses(")((()"), " - 2")
	fmt.Println(longestValidParentheses(")(()("), " - 2")
	fmt.Println(longestValidParentheses(")(())"), " - 4")
	fmt.Println(longestValidParentheses(")()(("), " - 2")
	fmt.Println(longestValidParentheses(")()()"), " - 4")
	fmt.Println(longestValidParentheses(")())("), " - 2")
	fmt.Println(longestValidParentheses(")()))"), " - 2")
	fmt.Println(longestValidParentheses("))((("), " - 0")
	fmt.Println(longestValidParentheses("))(()"), " - 2")
	fmt.Println(longestValidParentheses("))()("), " - 2")
	fmt.Println(longestValidParentheses("))())"), " - 2")
	fmt.Println(longestValidParentheses(")))(("), " - 0")
	fmt.Println(longestValidParentheses(")))()"), " - 2")
	fmt.Println(longestValidParentheses("))))("), " - 0")
	fmt.Println(longestValidParentheses(")))))"), " - 0")

	fmt.Println(longestValidParentheses("(()))(()())((()()))"), " - 14")
	fmt.Println(longestValidParentheses("(()"), " - 2")
	fmt.Println(longestValidParentheses(")()())"), " - 4")
	fmt.Println(longestValidParentheses("()()()(()()"), " - 6")
	fmt.Println(longestValidParentheses("()()()(()()()()"), " - 8")
	fmt.Println(longestValidParentheses("(()))())("), " - 4")
	fmt.Println(longestValidParentheses("(()(((()"), " - 2")
	fmt.Println(longestValidParentheses("()))))())"), " - 2")

	fmt.Println(longestValidParentheses("))(()())())())()(()))))"), " - 8")
}
