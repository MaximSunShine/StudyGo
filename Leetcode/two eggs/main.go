package main

import (
	"fmt"
	"math"
)

func FindFloor(secretfloor int, colFloor int) (int, int, int) {
	var eggs int = 2
	var jump int = int(math.Floor(math.Sqrt(float64(2 * colFloor))))
	var step int = 0
	min, max := 1, colFloor

	for i := jump; i <= colFloor; i += jump {
		step++
		if i > secretfloor {
			eggs--
			max = i
			for j := min + 1; j < max-1; j++ {
				step++
				if j > secretfloor {
					eggs--
					max = j - 1
					break
				}
			}
			break
		}
		jump--
		min = i
	}

	if eggs > 0 {
		step++
		if max > secretfloor {
			eggs--
			return max - 1, eggs, step
		}

	}
	
	return max, eggs, step
}

func main() {
	for secretFloor := 100; secretFloor > 1; secretFloor-- {
		fmt.Println(FindFloor(secretFloor, 100))
	}

}
