package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var a, b int16
	var col, i int16

	var array []int16

	r := bufio.NewReader(os.Stdin)

	fmt.Fscan(r, &col)
	for i = 0; i < col; i++ {
		fmt.Fscan(r, &a, &b)
		//fmt.Println(d)
		array = append(array, a+b)
	}

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	for _, v := range array {
		fmt.Fprintln(w, v)
	}

}
