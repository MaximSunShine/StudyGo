package main

import (
	"fmt"
	"strconv"
	"time"
)

func mySqrt3(x int) int {
	if x < 2 {
		return x
	}
	l, r := 1, x/2

	for l != r {
		a := float64(r*r) - float64(x)
		//b := float64(x) - float64(l*l)

		c := a / float64(x)
		if c == 0 {
			break
		}
		i := r
		r = l + int(float64(l)*c) //46340
		if i == r {
			if (l+1)*(l+1) <= x {
				return l + 1
			} else {
				return l
			}
		}
		//r = l + l - i
	}
	return l
}

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	len_x := len(strconv.FormatInt(int64(x), 2))
	len_rez := len_x/2 + len_x%2

	l := 1 << (len_rez - 1)
	r := 1 << (len_rez)

	for l != r {
		a := float64(r*r) - float64(l*l)
		b := float64(x) - float64(l*l)

		c := b / a
		if c == 0 {
			break
		}
		i := l
		l += int(float64(r-l) * c) //46340
		if i == l {
			if (l+1)*(l+1) <= x {
				return l + 1
			} else {
				return l
			}
		}
		r = l + l - i
	}
	return l
}

func help1(x, l, r int) (int, int) {

	a := float64(r*r) - float64(l*l)
	b := float64(x) - float64(l*l)

	c := b / a
	y := l + int(float64(r-l)*c) //46340
	return y, int(float64(y) * 1.07)
}
func mySqrt2(x int) int {
	if x < 2 {
		return x
	}
	l, r := 1, x/2

	if x > 1024 {
		len_x := len(strconv.FormatInt(int64(x), 2))
		len_rez := len_x/2 + len_x%2

		l = 1 << (len_rez - 1)
		r = 1<<(len_rez) - 1

		l, r = help1(x, l, r)
	}

	for i := l; l < r; i = l + (r-l)/2 + 1 {
		cur := i * i
		if cur < x {
			l = i
			continue
		}
		if cur > x {
			r = i - 1
			continue
		}
		if cur == x {
			return i
		}

		l = i
		break
	}
	return l
}

func main() {
	fmt.Println(mySqrt3(1681))
	//fmt.Println(mySqrt(0))
	//fmt.Println(mySqrt(1))
	//fmt.Println(mySqrt(2))
	//fmt.Println(mySqrt(3))
	//fmt.Println(mySqrt(4))
	//fmt.Println(mySqrt(8))
	//fmt.Println(mySqrt(9))
	//fmt.Println(mySqrt(10))
	//fmt.Println(mySqrt(15))
	//fmt.Println(mySqrt(63))
	//fmt.Println(mySqrt(65))
	//fmt.Println(mySqrt(1025))
	//fmt.Println(mySqrt(math.MaxInt32))

	t := time.Now()

	for i := 2147480647; i < 2147483647; /*2147483647*/ i++ {
		mySqrt(i)
	}
	fmt.Println("Done 1", time.Now().Nanosecond()-t.Nanosecond())

	t = time.Now()
	for i := 2147480647; i < 2147483647; /*2147483647*/ i++ {
		mySqrt2(i)
	}
	fmt.Println("Done 2", time.Now().Nanosecond()-t.Nanosecond())

}
