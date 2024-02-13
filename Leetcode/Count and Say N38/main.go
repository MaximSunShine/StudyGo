package main

import "fmt"

func countAndSay(n int) string {

	start := "1"

	b := []byte(start)
	var b_new []byte

	for i := 0; i < n-1; i++ {
		help := b[0]
		step := 0
		for j, v := range b {
			if v == help {
				step++
			} else {
				b_new = append(b_new, byte(step)+48, help)
				help = v
				step = 1
			}
			if j == len(b)-1 {
				b_new = append(b_new, byte(step)+48, help)
			}

		}
		b = b_new
		b_new = []byte{}
	}
	return string(b)

}
func main() {

	fmt.Println(countAndSay(1))
	fmt.Println(countAndSay(2))
	fmt.Println(countAndSay(3))
	fmt.Println(countAndSay(4))
	fmt.Println(countAndSay(5))
	fmt.Println(countAndSay(6))
	fmt.Println(countAndSay(7))
	fmt.Println(countAndSay(8))
	fmt.Println(countAndSay(9))
	fmt.Println(countAndSay(10))

}
