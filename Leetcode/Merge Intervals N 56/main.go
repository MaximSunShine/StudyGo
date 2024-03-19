package main

import (
	"fmt"
	"sort"
)

func insert1(aa [][]int, new []int) [][]int {
	aa = append(aa, new)
	if len(aa) < 2 {
		return aa
	}

	res := make([][]int, 0)
	sort.Slice(aa, func(a, b int) bool {
		return aa[a][0] < aa[b][0]
	})
	step := 0
	res = append(res, aa[0])
	for i := 1; i < len(aa); i++ {
		if res[step][1] >= aa[i][0] {
			if res[step][1] < aa[i][1] {
				res[step][1] = aa[i][1]
			}
		} else {
			res = append(res, aa[i])
			step++
		}
	}

	return res
}
func insert(intervals [][]int, newInterval []int) [][]int {
	// [[1,2],[3,5],[6,7],[8,10],[12,16]]      [4,8]
	//  3 10
	// if can't be merged add it to res
	// [1,2]
	n := len(intervals)
	i := sort.Search(n, func(i int) bool { return intervals[i][0] > newInterval[0] })
	j := sort.Search(n, func(j int) bool { return intervals[j][1] > newInterval[1] })
	if i >= 1 && newInterval[0] <= intervals[i-1][1] {
		newInterval[0] = intervals[i-1][0]
		i--
	}
	if j < n && newInterval[1] >= intervals[j][0] {
		newInterval[1] = intervals[j][1]
		j++
	}
	return append(intervals[:i], append([][]int{newInterval}, intervals[j:]...)...)
}

// 6,9
// 1,5

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(a, b int) bool { return intervals[a][0] < intervals[b][0] })

	for i := 0; i < len(intervals)-1; i++ {
		if intervals[i][1] >= intervals[i+1][0] {
			if intervals[i][1] < intervals[i+1][1] {
				intervals[i][1] = intervals[i+1][1]
			}

			intervals = append(intervals[:i+1], intervals[i+2:]...)
			i--
		}
	}

	return intervals

}
func main() {

	fmt.Println(merge([][]int{{1, 4}, {0, 2}, {3, 5}}))
	fmt.Println(insert([][]int{{1, 4}, {0, 2}, {3, 5}}, []int{2, 6}))

}
