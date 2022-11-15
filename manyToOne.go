package main

import (
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

func doTick(stop, echo chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	t := time.NewTicker(time.Millisecond * 200)
	for {
		select {
		case <-t.C:
			fmt.Println("do tick")
			f := func() {
				n := rand.Intn(1) + 1
				fmt.Println(n)
				if n == 1 {
					fmt.Println("Max is found")
					echo <- struct{}{}
				}
			}

			go f()
			go f()

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

	wg := new(sync.WaitGroup)

	stop := make(chan struct{})
	echo := make(chan struct{})
	wg.Add(2)
	go workTime(stop, echo, wg)
	go doTick(stop, echo, wg)
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
