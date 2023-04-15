package main

var start, last int

func expandWord1(start int, last int, a []byte) (int, int) {
	for start > 0 && last < len(a)-1 {
		if a[start-1] == a[last+1] {
			start--
			last++
		} else {
			break
		}
	}
	return start, last
}
func expandWord(a int, b int, arr []byte) {

	for a >= 0 && b < len(arr) && arr[a] == arr[b] {
		a--
		b++
	}

	if last-start < b-a-1 {
		last, start = b-1, a+1
	}

}
func longestPalindrome(w string) string {

	start, last = 0, 0
	arr := []byte(w)
	step := len(arr)/2 - 1

	for ; step < len(arr)-1 && last < len(arr)-1; step++ {
		expandWord(step, step, arr)
		expandWord(step, step+1, arr)
	}

	step = len(arr) / 2
	if last-start != len(arr)-1 {
		for ; step > 0; step-- {
			expandWord(step, step, arr)
			expandWord(step-1, step, arr)

		}
	}

	return string(arr[start : last+1])
}

func longestPalindrome1(w string) string {
	arr := []byte(w)

	var a, b, step, start, last int

	step = len(arr)/2 - 1

	for step < len(arr)-1 && b < len(arr)-1 {
		if arr[step] == arr[step+1] { // 100% palindromic
			a, b = expandWord1(step, step+1, arr)

			if b-a > last-start {
				start, last = a, b
			}
		}
		if step+2 < len(arr) {
			if arr[step] == arr[step+2] { // 100% palindromic
				a, b = expandWord1(step, step+2, arr)

				if b-a > last-start {
					start, last = a, b
				}
			}
		}
		step++
	}
	step = len(arr) / 2
	a = step
	if last-start != len(arr)-1 {
		for step > 0 && a > 0 {
			if arr[step] == arr[step-1] { // 100% palindromic
				a, b = expandWord1(step, step-1, arr)

				if b-a > last-start {
					start, last = a, b
				}
			}
			if step-2 >= 0 {
				if arr[step] == arr[step-2] { // 100% palindromic
					a, b = expandWord1(step, step-2, arr)

					if b-a > last-start {
						start, last = a, b
					}
				}
			}
			step--
		}
	}

	return string(arr[start : last+1])
}

func longestPalindrome2(b string) string {

	a := []byte(b)

	start := 0
	last := 0
	current := 0

	startMain := 0
	lastMain := 0

	for current < len(a)-1 && startMain != len(a)-1 { //дополнить если найденное слово больше остатка длины

		if a[current] == a[current+1] { // 100% palindromic
			start, last = expandWord1(current, current+1, a)

			if last-start > lastMain-startMain {
				startMain, lastMain = start, last
			}

		}
		if current+2 < len(a) {
			if a[current] == a[current+2] { //100% palindromic
				start, last = expandWord1(current, current+2, a)

				if last-start > lastMain-startMain {
					startMain, lastMain = start, last
				}

			}
		}

		current++

	}
	return string(a[startMain : lastMain+1])
}
func main() {
	println("-> ", longestPalindrome("qwerty"))
	println("-> ", longestPalindrome("aacabdkacaa"))
	println("-> ", longestPalindrome("adam"))
	println("-> ", longestPalindrome("ccc"))
	println("-> ", longestPalindrome("qwewdababagty"))
	println("-> ", longestPalindrome("qwewdabbag"))
}
