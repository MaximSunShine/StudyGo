package main

import "fmt"

func multiply(s1 string, s2 string) string {
	if s1[0] == '0' || s2[0] == '0' {
		return "0"
	}

	return convertToStr(mul(s1, s2))
}

func mul(s1, s2 string) []byte {
	var res = make([]byte, len(s1)+len(s2), len(s1)+len(s2))

	for i := len(s1) - 1; i >= 0; i-- {
		for j := len(s2) - 1; j >= 0; j-- {
			h := i + j
			res[h+1] += (s1[i] - '0') * (s2[j] - '0')
			if res[h+1] > 9 {
				res[h] += res[h+1] / 10
				res[h+1] = res[h+1] % 10
			}
		}
	}
	return res
}

func convertToStr(b []byte) string {
	if b[0] == 0 {
		b = b[1:] // clear first zero
	}

	for i := 0; i < len(b); i++ {
		b[i] += '0' // prepare to str
	}

	return string(b)
}

func main() {
	fmt.Println(multiply("123456789", "0"))
	fmt.Println(multiply("60974249908865105026646412538664653190280198809433017", "500238825698990292381312765074025160144624723742"))
	fmt.Println("30501687172287445993560048081057096686019986405658336826483685740920318317486606305094807387278589614" == multiply("60974249908865105026646412538664653190280198809433017", "500238825698990292381312765074025160144624723742"))
}

//30501687172287445993560048081057096686019986405658336826483685740920318317486606305094807387278589614
//30501687172287445993560048081057096686019986405658336826483685740920318317486606305094807387278589614
