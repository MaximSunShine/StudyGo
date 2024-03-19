package main

import "fmt"

func uniquePathsWithObstacles(aa [][]int) int {
	if aa[0][0] == 1 || aa[len(aa)-1][len(aa[0])-1] == 1 {
		return 0
	}

	dp := make([]int, len(aa[0]))
	dp[0] = 1
	for _, row := range aa {
		t := false
		for j, val := range row {
			if val == 1 {
				dp[j] = 0
			} else if j > 0 {
				dp[j] += dp[j-1]
			}
			if dp[j] > 0 {
				t = true
			}
		}
		if !t {
			return 0
		}
	}
	return dp[len(aa[0])-1]
}

func uniquePathsWithObstacles1(aa [][]int) int {
	if aa[0][0] == 1 || aa[len(aa)-1][len(aa[0])-1] == 1 {
		return 0
	}

	for i := 0; i < len(aa); i++ {
		t := false
		for j := 0; j < len(aa[0]); j++ {

			if aa[i][j] == 0 {
				if i+j == 0 {
					aa[i][j] = 1
					t = true
					continue
				}
				if i-1 >= 0 {
					aa[i][j] = aa[i-1][j]
				}
				if j-1 >= 0 {
					aa[i][j] += aa[i][j-1]
				}
				if aa[i][j] > 0 {
					t = true
				}
			} else {
				aa[i][j] = 0
			}
		}
		if !t {
			return 0
		}
	}

	return aa[len(aa)-1][len(aa[0])-1]
}

func main() {
	fmt.Println(uniquePathsWithObstacles([][]int{{1, 1}}))                          //0
	fmt.Println(uniquePathsWithObstacles([][]int{{1, 0}}))                          //0
	fmt.Println(uniquePathsWithObstacles([][]int{{0, 0}, {0, 1}}))                  //0
	fmt.Println(uniquePathsWithObstacles([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}})) //2
	fmt.Println(uniquePathsWithObstacles([][]int{{0, 1}, {0, 0}}))                  //1
	fmt.Println(uniquePathsWithObstacles([][]int{{1}, {0}}))                        //0
}
