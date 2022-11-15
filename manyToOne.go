package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func workTime(stop, echo chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	t := time.NewTimer(time.Second * 1)
	select {
	case <-t.C:
		fmt.Println("Timer is stoped")
		echo <- struct{}{}
	case <-stop:
		fmt.Println("Close the Timer")
		t.Stop()
	}
}

func doTick(stop, echo chan struct{}, wg *sync.WaitGroup, num int) {
	defer wg.Done()
	t := time.NewTicker(time.Millisecond * 199)
	for {
		select {
		case <-t.C:
			fmt.Println("do tick")

			wg1 := new(sync.WaitGroup)
			f := func() {
				defer wg1.Done()
				n := rand.Intn(num)
				fmt.Println(n)
				if n == 0 {
					fmt.Println("Max is found")
					echo <- struct{}{}
				}
			}

			wg1.Add(3)
			go f()
			go f()
			go f()
			wg1.Wait()

		case <-stop:
			fmt.Println("stop Ticker")
			t.Stop()
			return
		}
	}
}

func listen(stop, echo chan struct{}) {
	b := true
	for range echo {
		if b {
			fmt.Println("Stop all gorutines!!!")
			close(stop)
			b = false
		}
	}
}

func main() {
	num := flag.Int("n", 2, "Предел случайных чисел от 0 до n-1")
	flag.Parse()
	wg := new(sync.WaitGroup)

	stop := make(chan struct{})
	echo := make(chan struct{})
	wg.Add(2)
	go workTime(stop, echo, wg)
	go doTick(stop, echo, wg, *num)
	go listen(stop, echo)

	ch := make(chan struct{})
	go func() {
		defer close(echo)
		defer close(ch)
		wg.Wait()
	}()
	<-ch
	fmt.Println("Finish")
}
