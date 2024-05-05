package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	t := 0
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		bank := make([][]float64, 3) //!!

		for j := 0; j < 3; j++ {
			for k := 0; k < 6; k++ {
				a, b := 0.0, 0.0
				fmt.Fscan(in, &a, &b)
				bank[j] = append(bank[j], b/a)
			}
		}

		rez := make([]float64, 15)

		rez = append(rez, bank[0][0]) //p/d
		rez = append(rez, bank[1][0])
		rez = append(rez, bank[2][0])

		rez = append(rez, bank[0][1]*bank[1][5]) //e/d
		rez = append(rez, bank[0][1]*bank[2][5])
		rez = append(rez, bank[1][1]*bank[0][5])
		rez = append(rez, bank[1][1]*bank[2][5])
		rez = append(rez, bank[2][1]*bank[0][5])
		rez = append(rez, bank[2][1]*bank[1][5])

		rez = append(rez, bank[0][0]*bank[1][3]*bank[2][5]) //d/e/d
		rez = append(rez, bank[0][0]*bank[2][3]*bank[1][5])
		rez = append(rez, bank[1][0]*bank[0][3]*bank[2][5])
		rez = append(rez, bank[1][0]*bank[2][3]*bank[0][5])
		rez = append(rez, bank[2][0]*bank[0][3]*bank[1][5])
		rez = append(rez, bank[2][0]*bank[1][3]*bank[0][5])

		rez = append(rez, bank[0][0]*bank[1][2]*bank[2][0]) //d/r/d
		rez = append(rez, bank[0][0]*bank[2][2]*bank[1][0])
		rez = append(rez, bank[1][0]*bank[0][2]*bank[2][0])
		rez = append(rez, bank[1][0]*bank[2][2]*bank[0][0])
		rez = append(rez, bank[2][0]*bank[0][2]*bank[1][0])
		rez = append(rez, bank[2][0]*bank[1][2]*bank[0][0])

		rez = append(rez, bank[0][1]*bank[1][4]*bank[2][0]) //e/r/d
		rez = append(rez, bank[0][1]*bank[2][4]*bank[1][0])
		rez = append(rez, bank[1][1]*bank[0][4]*bank[2][0])
		rez = append(rez, bank[1][1]*bank[2][4]*bank[0][0])
		rez = append(rez, bank[2][1]*bank[0][4]*bank[1][0])
		rez = append(rez, bank[2][1]*bank[1][4]*bank[0][0])

		max := 0.0
		for _, v := range rez {
			if max < v {
				max = v
			}
		}
		fmt.Fprintln(out, max)
	}
}
