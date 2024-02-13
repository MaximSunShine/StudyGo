package main

func rotate(matrix [][]int) {

	if len(matrix) < 2 {
		return
	}

	f := func(m [][]int) {

		for step := 0; step*step < (len(m) - 1); step++ {
			for i := step; i < len(m)-1-step; i++ {
				help := m[step][i]
				m[step][i] = m[len(m)-1-i][step]
				m[len(m)-1-i][step] = m[len(m)-1-step][len(m)-1-i]
				m[len(m)-1-step][len(m)-1-i] = m[i][len(m)-1-step]
				m[i][len(m)-1-step] = help
			}
		}
	}
	f(matrix)
}

func main() {
	rotate([][]int{})
	rotate([][]int{{1}})
	rotate([][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}, {21, 22, 23, 24, 25}})

}
