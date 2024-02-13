package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file1, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file1.Close()

	in := bufio.NewReader(file)
	out := bufio.NewWriter(file1)
	defer out.Flush()

	a := 0
	fmt.Fscan(in, &a)
	b := 0
	fmt.Fscan(in, &b)

	fmt.Fprintln(out, a+b)
}
