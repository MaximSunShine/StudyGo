/*package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	n1      = flag.Int("N1", 10, "Этап №-1 кол-во горутин")
	m1      = flag.Int("M1", 7, "Этап №-2 кол-во горутин")
	k1      = flag.Int("K1", 15, "Этап №-3 кол-во горутин")
	n2      = flag.Int("N2", 200, "Интервал N2 в млсек")
	a       = flag.Int("a", 4, "Начальное значение интервала случайного числа")
	b       = flag.Int("b", 9, "Конечное значение интервала случайного числа")
	allTime = flag.Int("B", 2, "Время работы программы в сек")
)

func allDiv(n int) []int {
	var a = []int{1}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			a = append(a, i)
		}
	}
	a = append(a, n)
	return a
}

func findNum(b [2]int, wg *sync.WaitGroup, num_ch chan int, findMax_ch chan struct{}) {
	defer wg.Done()
	v := rand.Intn(b[1]-b[0]+1) + b[0]
	if v == b[1] {
		fmt.Println("max element done")
		findMax_ch <- struct{}{} // отправляем в канал метку, что максимальный элемент был найден
	}
	num_ch <- v
}

func makeTick(t *time.Ticker, stop chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-stop:
			fmt.Println("Tick stopped")
			return
		case <-t.C: // возможно через дефалт
			fmt.Println("tick do")
		}
	}
}

func makeTimeEnd(t *time.Ticker, echo, stop chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	t_all := time.NewTimer(time.Second * time.Duration(*allTime))
	defer t.Stop()
	select {
	case <-t_all.C:
		fmt.Println("End Timer")
		echo <- struct{}{} // оповещаем о приостановке всех горутин
	case <-stop:
		fmt.Println("Timer stopped")
		t_all.Stop()
	}
}

func main() {
	t := time.NewTicker(time.Millisecond * time.Duration(*n2))
	stop := make(chan struct{})
	echo := make(chan struct{})

	flag.Parse()
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go makeTick(t, stop, wg)
	go makeTimeEnd(t, echo, stop, wg)

	go func() { // слушаем если кто-то оповестил о приостановке работ горутин
		i := true
		for range echo {
			if i {
				fmt.Println("Make stopping goroutine")
				close(stop) // останавливаем все горутины
				i = false
			}
		}
		//fmt.Println("echo closed")
	}()

	go func() { // оповещение об остановке всех горутин через тест время
		t := time.After(time.Second * 1)
		<-t
		fmt.Println("Say stop the goroutine")
		echo <- struct{}{}
	}()
	wg.Wait()
	close(echo)
	//t1 := time.NewTimer(time.Second * 1)
	//<-t1.C
	fmt.Println("Finish")
}
