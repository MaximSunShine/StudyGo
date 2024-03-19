package main

import "fmt"

func minDistance(word1 string, word2 string) int {
	l1 := len(word1)
	l2 := len(word2)
	cache := make([][]int, l1+1)
	// Init base cases
	for i := range cache {
		cache[i] = make([]int, l2+1)
		cache[i][l2] = l1 - i
	}
	cache[l1] = make([]int, l2+1)
	for i := range cache[l1] {
		cache[l1][i] = l2 - i
	}

	for i := l1 - 1; i >= 0; i-- {
		for j := l2 - 1; j >= 0; j-- {
			if word1[i] == word2[j] {
				cache[i][j] = cache[i+1][j+1]
			} else {
				insert := cache[i][j+1]
				delete := cache[i+1][j]
				replace := cache[i+1][j+1]
				cache[i][j] = 1 + min(min(insert, delete), replace)
			}
		}
	}

	return cache[0][0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {

	fmt.Println(minDistance("horse", "ros"))
	fmt.Println(minDistance("intention", "execution"))

}
