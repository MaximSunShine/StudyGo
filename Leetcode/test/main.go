package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	data           = []DataCenter{}
	colDataCenters int
	colServers     int
	colActions     int
)

type DataCenter struct {
	servers     []bool
	colWorks    int
	colRestarts int
}

func getMin() int {
	min := colActions * colServers
	imin := 0
	for i, val := range data {
		/*col := 0
		for _, v := range val.servers {
			if !v {
				col++
			}
		}*/
		help := data[i].colWorks * val.colRestarts
		if min > help {
			min = help
			imin = i
		}
	}

	return imin + 1
}

func getMax() int {
	imax := 0
	max := 0
	for i, val := range data {
		/*col := 0
		for _, v := range val.servers {
			if !v {
				col++
			}
		}*/
		help := data[i].colWorks * val.colRestarts
		if max < help {
			max = help
			imax = i
		}
	}

	return imax + 1
}

func disable(i int, j int) {
	if !data[i].servers[j] {
		data[i].servers[j] = true
		data[i].colWorks--
		//data[i].colRestarts++
	}
}
func reset(i int) {

	for j, _ := range data[i].servers { //val do not
		data[i].servers[j] = false //val do not
	}
	data[i].colWorks = colServers

	data[i].colRestarts++
}

func main() {

	//ss := "3 3 12\nDISABLE 1 2\nDISABLE 2 1\nDISABLE 3 3\nGETMAX\nRESET 1\nRESET 2\nDISABLE 1 2\nDISABLE 1 3\nDISABLE 2 2\nGETMAX\nRESET 3\nGETMIN"

	fmt.Scanln(&colDataCenters, &colServers, &colActions)

	data = make([]DataCenter, colDataCenters, colDataCenters)

	for i := 0; i < colDataCenters; i++ {
		data[i] = DataCenter{
			servers:     make([]bool, colServers, colServers),
			colWorks:    colServers,
			colRestarts: 0,
		}
	}

	var s string
	scanner := bufio.NewScanner(os.Stdin)

	//fmt.Scanf("%s\n", &s)
	for i := 0; i < colActions; i++ {

		scanner.Scan()
		s = scanner.Text()
		a := strings.Split(s, " ")

		if string(a[0]) == "DISABLE" {
			x, _ := strconv.Atoi(a[1]) //проверка на ошибку!
			y, _ := strconv.Atoi(a[2])
			disable(x-1, y-1)
		} else if string(a[0]) == "RESET" {
			x, _ := strconv.Atoi(a[1])
			reset(x - 1)
		} else if string(a[0]) == "GETMAX" {
			fmt.Println(getMax())
		} else if string(a[0]) == "GETMIN" { // getMin
			fmt.Println(getMin())
		}

	}

}
