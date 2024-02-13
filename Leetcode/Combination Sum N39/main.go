package main

import (
	"fmt"
	"sort"
	"sync"
)

func recursion(ch chan<- []int, candidates []int, answer []int, sum int, step int, target int, bottom int, wg *sync.WaitGroup) {

	//fmt.Println(bottom)
	//mu := new(sync.Mutex)
	defer wg.Done()
	for i := step; i < len(candidates); i++ {
		if sum+candidates[i] == target {
			answer = append(answer, candidates[i])
			help := []int{}
			for _, v := range answer {
				help = append(help, v)
			}
			ch <- help
			//fmt.Println(answer)
			return
		} else if sum+candidates[i] < target {
			sum += candidates[i]
			answer = append(answer, candidates[i])
		} else {
			return
		}
		wg.Add(1)
		recursion(ch, candidates, answer, sum, i, target, bottom+1, wg)
		answer = answer[0 : len(answer)-1]
		sum -= candidates[i]
		//fmt.Println(bottom)
	}
}

func combinationSum(candidates []int, target int) [][]int {

	sort.Ints(candidates)
	ch := make(chan []int)
	var result [][]int
	stop := make(chan struct{})

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go recursion(ch, candidates, []int{}, 0, 0, target, 0, wg)

	go func() {
		//mu := new(sync.Mutex)
		for v := range ch {
			result = append(result, v)
		}
		close(stop)
	}()
	go func() {
		wg.Wait()
		close(ch)

	}()

	<-stop
	return result
}
func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
	fmt.Println(combinationSum([]int{2}, 1))
	fmt.Println(combinationSum([]int{8, 7, 4, 3}, 11))

}
