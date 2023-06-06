package main

import "fmt"

func letterCombinations(digits string) []string {
	a := []string{
		0: "",
		1: "",
		2: "abc",
		3: "def",
		4: "ghi",
		5: "jkl",
		6: "mno",
		7: "pqrs",
		8: "tuv",
		9: "wxyz",
	}

	b := []int{1, 1, 1, 1}

	for i, v := range digits {
		b[i] = int(v) - 48
	}

	C := make([]string, 0, b[0]*b[1]*b[2]*b[3])

	for i := 0; i < len(a[b[0]]); i++ {
		A := string(a[b[0]][i])

		for j := 0; j < len(a[b[1]]); j++ {
			B := A + string(a[b[1]][j])

			for k := 0; k < len(a[b[2]]); k++ {
				K := B + string(a[b[2]][k])

				for l := 0; l < len(a[b[3]]); l++ {
					L := K + string(a[b[3]][l])
					C = append(C, L)
				}

				if b[3] == 1 {
					C = append(C, K)
				}
			}
			if b[2] == 1 {
				C = append(C, B)
			}
		}
		if b[1] == 1 {
			C = append(C, A)
		}
	}

	return C
}

func main() {
	fmt.Println(letterCombinations("5678"))
}
