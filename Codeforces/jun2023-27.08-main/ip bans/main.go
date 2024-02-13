package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("C:/Users/Zver/Downloads/tests (1)/tests/01")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file1, err := os.Create("C:/Users/Zver/Downloads/tests (1)/tests/01.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file1.Close()

	in := bufio.NewReader(file)
	out := bufio.NewWriter(file1)
	defer out.Flush()
	n, k := 0, 0

	a, z := fmt.Fscan(in, &n, &k)
	fmt.Fprintln(out, a, z)
	s := ""
	var arr [][]string

	hash := make(map[string]int)

	sum := 0

	stop := false
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &s)
		arr = append(arr, strings.Split(s, "."))

		hash[arr[i][2]]++
		if hash[arr[i][2]] == 256 {
			stop = true
		}
		if hash[arr[i][2]] == 2 {
			sum += 254
		} else if hash[arr[i][2]] > 2 {
			sum--
		}
	}

	//fmt.Println()
	if len(hash) == 65536 {
		fmt.Fprintln(out, 65534)
		fmt.Fprintln(out, 1)
		fmt.Fprintln(out, "100.200.0.0/16")
	} else {
		if len(hash) > k { //------------/16
			fmt.Fprintln(out, 65536-len(arr))
			fmt.Fprintln(out, 1)
			fmt.Fprintln(out, "100.200.0.0/16")

		} else if len(arr) <= k && !stop { //----------0.0.0.0 проработать ели все комбинации по маске выбраны, минимизировать
			fmt.Fprintln(out, 0)
			fmt.Fprintln(out, len(arr))

			for i := 0; i < n; i++ {
				help := "100.200." + arr[i][2] + "." + arr[i][3]
				fmt.Fprintln(out, help)
			}

			//fmt.Fprintln(out, "...")

		} else if len(hash) <= k { //-----------/24

			fmt.Fprintln(out, sum)
			fmt.Fprintln(out, len(hash))
			for key, val := range hash {
				if val > 1 {
					help := "100.200." + key + ".0/24"
					fmt.Fprintln(out, help)
				} else {
					for i := 0; i < n; i++ {
						if arr[i][2] == key {
							help := "100.200." + key + "." + arr[i][3]
							fmt.Fprintln(out, help)
						}
					}
				}

			}
		}
	}

	//fmt.Fprintln(out, hash)
}
