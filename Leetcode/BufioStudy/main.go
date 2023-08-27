package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	sc.Scan()
	N, _ := strconv.Atoi(sc.Text())

	c := make([]int, N)
	for i := 0; i < N; i++ {
		sc.Scan()
		c[i], _ = strconv.Atoi(sc.Text())
	}

	hash := make(map[int]int)

	for i := 0; i < N; i++ {
		sc.Scan()
		hash[c[i]], _ = strconv.Atoi(sc.Text())
	}
	sc.Scan()
	K, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	s, _ := strconv.Atoi(sc.Text())

	cur := 0
	if K > 0 {
		cur = hash[s]
	}

	step := 0
	for i := 1; i < K; i++ {
		sc.Scan()
		s, _ := strconv.Atoi(sc.Text())
		if hash[s] != cur {
			cur = hash[s]
			step++
		}
	}
	fmt.Print(step)
}

/*s = bufio.NewScanner(os.Stdin)
s.Scan()
//s.Split(bufio.ScanWords)

//ss := strings.NewReader("Привет Макс")

for s.Scan() {
	err := s.Err()
	if err == nil {
		log.Println("Scan completed and reached EOF")
	} else {
		log.Fatal(err)
	}

	fmt.Println(s.Text())
}
*/
