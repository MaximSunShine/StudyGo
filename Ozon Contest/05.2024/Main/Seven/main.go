package main

import (
	"bufio"
	"cmp"
	"fmt"
	"log"
	"os"
	"slices"
	"time"
)

type Arrival struct {
	timeIn    int
	carNumber int
}

type Car struct {
	start     int
	end       int
	capacity  int
	carNumber int
}

func main() {
	dir := ".\\tests\\"
	list, _ := os.ReadDir(dir)

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	t := time.Now()
	t1 := t
	var t2 time.Time
	for numTest := 0; numTest < len(list); numTest += 2 {
		dataTest, answerTest := list[numTest].Name(), list[numTest+1].Name()

		fileReadData, err := os.Open(dir + dataTest)
		if err != nil {
			log.Fatal(err)
		}

		fileReadAnswer, err := os.Open(dir + answerTest)
		if err != nil {
			log.Fatal(err)
		}
		//your code is here
		in := bufio.NewReader(fileReadData)
		tests := bufio.NewReader(fileReadAnswer)

		numSets := 0
		fmt.Fscan(in, &numSets)

		for num := 0; num < numSets; num++ {
			colArr := 0
			fmt.Fscan(in, &colArr)
			arrivals := make([]Arrival, colArr, colArr)
			sortArrivals := make([]*Arrival, colArr, colArr)
			for i := 0; i < colArr; i++ {
				fmt.Fscan(in, &arrivals[i].timeIn)
				sortArrivals[i] = &arrivals[i]
			}

			colCars := 0
			fmt.Fscan(in, &colCars)
			cars := make([]Car, colCars, colCars)
			carsStartTime := make([]*Car, colCars, colCars)
			for i := 0; i < colCars; i++ {
				fmt.Fscan(in, &cars[i].start, &cars[i].end, &cars[i].capacity)
				cars[i].carNumber = i
				carsStartTime[i] = &cars[i]
			}
			ch := make(chan struct{})
			go func() {
				slices.SortFunc(sortArrivals, func(i, j *Arrival) int {
					return cmp.Compare(i.timeIn, j.timeIn)
				})
				ch <- struct{}{}
			}()

			slices.SortStableFunc(carsStartTime, func(a, b *Car) int {
				// If starts are equal, order by carNumber
				return cmp.Compare(a.start, b.start)
			})
			<-ch

			/*slices.SortFunc(carsStartTime, func(a, b *Car) int {
				if n := cmp.Compare(a.start, b.start); n != 0 {
					return n
				}
				// If starts are equal, order by carNumber     return cmp.Compare(a.Age, b.Age))
				return cmp.Compare(a.carNumber, b.carNumber)
			})*/

			for i, val := range sortArrivals {
				for len(carsStartTime) > 0 {
					if carsStartTime[0].capacity == 0 || carsStartTime[0].end < val.timeIn {
						carsStartTime = carsStartTime[1:]
						continue
					} else if val.timeIn >= carsStartTime[0].start {
						carsStartTime[0].capacity--
						sortArrivals[i].carNumber = carsStartTime[0].carNumber + 1
						break
					} else {
						sortArrivals[i].carNumber = -1
						break
					}
				}
				if len(carsStartTime) == 0 {
					val.carNumber = -1
				}
			}
			//s, _, _ := tests.ReadLine()
			for i, v := range arrivals {
				help := 0
				fmt.Fscan(tests, &help)
				if v.carNumber != help {
					fmt.Println("test N=", dataTest, "i=", i, "is=", v.carNumber, "need=", help)
					return
				}
			}
		}
		//must be in the end
		t2 = time.Now()
		fmt.Println("test N =", dataTest, "is done", t2.Sub(t1).Seconds(), "sec")
		t1 = t2

		fileReadData.Close()
		fileReadAnswer.Close()
	}
	fmt.Fprintln(out, "Total time is", t2.Sub(t))
	fmt.Fprintln(out, "Middle time is", t2.Sub(t)/time.Duration(len(list)))
	fmt.Scan()
}
