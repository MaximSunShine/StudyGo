package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type divider struct {
	N int
	A []int
}

var (
	n1      = flag.Int("N1", 10, "Этап №-1 кол-во горутин")
	m1      = flag.Int("M1", 7, "Этап №-2 кол-во горутин")
	k1      = flag.Int("K1", 15, "Этап №-3 кол-во горутин")
	n2      = flag.Int("N2", 200, "Интервал N2 в млсек")
	a       = flag.Int("a", 4, "Начальное значение интервала случайного числа")
	b       = flag.Int("b", 9, "Конечное значение интервала случайного числа")
	allTime = flag.Int("B", 2, "Время работы программы в сек")

	num_ch = make(chan int)
	stop   = make(chan struct{})
	echo   = make(chan struct{})
	div_ch = make(chan []byte)

	json_array = [][]byte{}
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

func findNum(b [2]int, wg *sync.WaitGroup) {
	defer wg.Done()
	v := rand.Intn(b[1]-b[0]+1) + b[0]
	if v == b[1] {
		fmt.Println("max element find")
		echo <- struct{}{} // отправляем в канал метку, что максимальный элемент был найден
	}
	num_ch <- v
}

func workTime(wg *sync.WaitGroup) {
	defer wg.Done()
	t := time.NewTimer(time.Second * time.Duration(*allTime))
	select {
	case <-t.C:
		fmt.Println("Timer is stopped")
		echo <- struct{}{}
	case <-stop:
		fmt.Println("Close the Timer")
		t.Stop()
	}
}

func doTick(wg *sync.WaitGroup) {
	defer wg.Done()
	t := time.NewTicker(time.Millisecond * time.Duration(*n2))

	go func() {
		defer close(div_ch)
		wg := new(sync.WaitGroup)
		for i := 0; i < *m1; i++ {
			wg.Add(1)
			go Find_dividers(wg)
		}
		wg.Wait()
	}()

	wg1 := new(sync.WaitGroup)
	do := true
	for do {
		select {
		case <-t.C:
			fmt.Println("do tick")
			for i := 0; i < *n1; i++ {
				wg1.Add(1)
				go findNum([2]int{*a, *b}, wg1)
			}

		case <-stop:
			fmt.Println("stop Ticker")
			t.Stop()
			do = false
		}
	}

	wg1.Wait()
	close(num_ch)
}

func listenEcho() {
	b := true
	for range echo {
		if b {
			fmt.Println("Stopping all goroutines!!!")
			close(stop)
			b = false
		}
	}
}

func Find_dividers(wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range num_ch { //???
		val := new(divider)
		val.N = v
		val.A = allDiv(v)

		data, err := json.MarshalIndent(val, "", "\t")
		if err != nil {
			fmt.Println(err)
			//return  ??? в канал ничего не пишем
		} else {
			div_ch <- data
		}
	}
}

func findMax() int {
	data := []divider{}
	max := 0
	for _, v := range json_array {
		var d divider
		if err := json.Unmarshal(v, &d); err != nil {
			//return???
		}
		data = append(data, d)
		if max < d.N {
			max = d.N
		}
	}
	return max
}

func main() {
	flag.Parse()
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go workTime(wg)
	go doTick(wg)

	go listenEcho()

	go func() {
		for v := range div_ch { // получаем из канала общие делители числа num
			json_array = append(json_array, v) // записываем даные: все делители числа num в массив
		}
	}()

	ch := make(chan struct{})
	go func() {
		defer close(echo)
		defer close(ch)
		//defer close(num_ch)
		wg.Wait()
	}()

	<-ch
	fmt.Println("Finish max number is ", findMax())
}
