package main

import (
	"fmt"
	"slices"
	"sync"
)

func exist(board [][]byte, word string) bool {
	r := []rune(word)
	slices.Reverse(r)
	word = string(r)

	if len(word) > len(board)*len(board[0]) {
		return false
	}

	var a = make([][]int, 0, len(board)*len(board[0]))
	var ch = make(chan []int)
	var stop = make(chan struct{})

	go func(stop <-chan struct{}, ch chan<- []int) {
		defer close(ch)

		for i, v := range board {
			for j, val := range v {
				select {
				case <-stop:
					return
				default:
					if val == word[0] {
						ch <- []int{i, j}
						a = append(a, []int{i, j})
					}
				}
			}
		}
	}(stop, ch)

	chRez := make(chan bool)

	f := func(stop chan struct{}, chRez chan<- bool, a []int, wg *sync.WaitGroup) {
		defer wg.Done()

		if len(word) == 1 {
			chRez <- true
			return
		}

		i, j := a[0], a[1]
		m := make(map[int]struct{})
		m[i*10+j] = struct{}{}
		step := 0

		var find func(int, int, int) bool
		find = func(i, j, step int) bool {
			//left
			key := i*10 + j - 1
			_, is := m[key]
			if !is && j > 0 && board[i][j-1] == word[step+1] {
				if step+1 == len(word)-1 {
					chRez <- true
					return true
				}
				m[key] = struct{}{}
				if find(i, j-1, step+1) {
					return true
				}
				delete(m, key)
			}
			//up
			key = (i-1)*10 + j
			_, is = m[key]
			if !is && i > 0 && board[i-1][j] == word[step+1] {
				if step+1 == len(word)-1 {
					chRez <- true
					return true
				}
				m[key] = struct{}{}
				if find(i-1, j, step+1) {
					return true
				}
				delete(m, key)
			}
			//right
			key = (i)*10 + j + 1
			_, is = m[key]
			if !is && j < len(board[0])-1 && board[i][j+1] == word[step+1] {
				if step+1 == len(word)-1 {
					chRez <- true
					return true
				}
				m[key] = struct{}{}
				if find(i, j+1, step+1) {
					return true
				}
				delete(m, key)
			}
			//down
			key = (i+1)*10 + j
			_, is = m[key]
			if !is && i < len(board)-1 && board[i+1][j] == word[step+1] {
				if step+1 == len(word)-1 {
					chRez <- true
					return true
				}
				m[key] = struct{}{}
				if find(i+1, j, step+1) {
					return true
				}
				delete(m, key)
			}

			return false

		}

		find(i, j, step)
	}

	var wg = new(sync.WaitGroup)
	wg.Add(1)
	go func(stop chan struct{}, ch <-chan []int) {
		defer wg.Done()

		for a := range ch {
			select {
			case <-stop:
				return
			default:
				wg.Add(1)
				go f(stop, chRez, a, wg)

			}
		}
	}(stop, ch)

	go func(wg *sync.WaitGroup) {

		wg.Wait()
		close(chRez)

	}(wg)

	rez := false
	oneTime := false
	for val := range chRez {
		if !oneTime {
			rez = val
			close(stop)
			oneTime = true
		}
	}

	return rez
}
func main() {
	fmt.Println(exist([][]byte{
		{'A', 'A', 'A', 'A', 'A', 'A'},
		{'A', 'A', 'A', 'A', 'A', 'A'},
		{'A', 'A', 'A', 'A', 'A', 'A'},
		{'A', 'A', 'A', 'A', 'A', 'A'},
		{'A', 'A', 'A', 'A', 'B', 'A'},
		{'A', 'A', 'A', 'A', 'B', 'A'}}, "AAAAAAAAAAAAABB"))
	fmt.Println(exist([][]byte{
		{'A'}}, "AA"))
	fmt.Println(exist([][]byte{
		{'A'}}, "B"))
	fmt.Println(exist([][]byte{
		{'A'}}, "A"))
	fmt.Println(exist([][]byte{
		{'A', 'A'}}, "A"))
	fmt.Println(exist([][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'}}, "ABCCED"))
	fmt.Println(exist([][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'}}, "SEE"))
	fmt.Println(exist([][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'}}, "ABCB"))

}
