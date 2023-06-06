package main

import "fmt"

func findMap(n int) map[string]struct{} {

	if n > 1 {
		m := findMap(n - 1)
		M := make(map[string]struct{})
		a := "()"
		for s := range m {
			for j := 0; j < len(s); j++ {
				M[s[:j]+a+s[j:]] = struct{}{}

			}
			//rez = append(rez, s+a)
			M[s+a] = struct{}{}
		}
		return M

	}
	return map[string]struct{}{
		"()": {},
	}

}

func generateParenthesis(n int) []string {
	if n == 1 {
		return []string{"()"}
	}

	help := findMap(n)

	var rez []string
	for i := range help {
		rez = append(rez, i)
	}

	return rez
}

func main() {

	fmt.Println(generateParenthesis(4))

}
