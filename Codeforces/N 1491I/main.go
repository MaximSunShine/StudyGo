package main

import "fmt"

type animal struct {
	num  int
	a    [3]int
	next *animal
}

func findKing(current *animal, colAnimal int) (int, int) {

	colBattle := 0
	count := 0
	pow := 0
	var help *animal
	for count < colAnimal && colBattle < 300000000 {
		colBattle++
		// second goes to the end of queue
		if current.next.a[pow] > current.next.next.a[0] {
			pow++
			count = 0

			help = current.next
			current.next = current.next.next
			help.next = current.next.next
			current.next.next = help
			current = current.next
			/*
				current.next, current.next.next, current.next.next.next = current.next.next, current.next.next.next, current.next
				current = current.next
			*/
			if pow == 3 {
				return current.next.num, colBattle
			}
		} else { // first goes to the end of queue
			count++
			current = current.next
			pow = 1
		}
	}

	return -1, -1
}

func main() {
	var col_animal int
	fmt.Scan(&col_animal)
	var first, current *animal

	current = &animal{}
	fmt.Scan(&current.a[0], &current.a[1], &current.a[2])
	current.num = 0
	first = current

	for i := 1; i < col_animal; i++ {
		help := &animal{}
		fmt.Scan(&help.a[0], &help.a[1], &help.a[2])
		help.num = i
		current.next = help
		current = help
	}

	current.next = first

	fmt.Println(findKing(current, col_animal))
}
