package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	allsec              int = 10
	maxtimefindProducts int = 5
	maxcolCustQ         int = 2
	maxtimeQ            int = 5
	maxcolCustEnter         = 10
)

var (
	chsec  [allsec + maxtimeQ]chan struct{} //событие для постановки в очередь в кассу
	chsecQ [allsec + maxtimeQ]chan struct{} //событие для выхода с кассы

	maxcolLiveCust int

	mu          = sync.Mutex{}
	colCust     int
	colLiveCust int
	queue       [10]int
	maxcashDesk int
)

func Queue(sec int, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	var qcurrent int
	r := rand.Intn(maxtimeQ) + 1
	secOut := sec + r //время покинуть магазин
	find := false
	for i := range queue { //поиск приемлемой кассы
		mu.Lock()
		if queue[i] < maxcolCustQ { //в кассе меньше N человек
			find = true
			queue[i]++
			qcurrent = i
			if maxcashDesk < i {
				maxcashDesk = i
			}
			fmt.Println(id, " Покупатель", queue[i], "-й в очереди на кассе N-", qcurrent)
			mu.Unlock() //необходимость
			break
		}
		mu.Unlock()
	}
	if !find {
		fmt.Println("Покупателю ", id, " Касс не хватает!!!")
	}
	select {
	case <-chsecQ[secOut]:
		mu.Lock()
		queue[qcurrent]--
		mu.Unlock()
		fmt.Println(id, " Покупатель покинул магазин")
	}
	mu.Lock()
	colLiveCust--
	mu.Unlock()
}

func newCustomer(sec int, colCustwg *sync.WaitGroup) {
	defer colCustwg.Done()
	r := rand.Intn(maxtimefindProducts) + 1 //время на поиск товаров
	secFindProducts := r + sec              //время пойти к кассе
	mu.Lock()
	colCust++     //+1 покупатель в магазине
	id := colCust //номер покупателя
	mu.Unlock()
	fmt.Printf("Покупатель N-%v хочет зайти в магазин в %v сек, нужно %v сек на покупки\n", id, sec, r)
	enter := false
	if secFindProducts < allsec { // успеет ли покупатель до закрытия магазина
		enter = true
		mu.Lock()
		colLiveCust++
		if maxcolLiveCust < colLiveCust {
			maxcolLiveCust = colLiveCust
		}
		mu.Unlock()

		select {
		case <-chsec[secFindProducts]: //прослушиваем канал (эхо) отвечающий за текущую секунду
			fmt.Printf("N-%v подошел к кассе в %v сек\n", id, secFindProducts)

			qwg := sync.WaitGroup{}
			qwg.Add(1)
			go Queue(secFindProducts, id, &qwg)
			qwg.Wait()

			return
		}
	}
	if !enter {
		fmt.Printf("!!!Покупатель N-%v не успел в магазин\n", id)
	}
}

func main() {

	for i := range chsec { //иннициализация каналов (эхо) всех сек работы магазина
		chsec[i] = make(chan struct{})  //не уверен, что надо
		chsecQ[i] = make(chan struct{}) //...
	}

	colCustwg := sync.WaitGroup{}
	for i := 0; i < allsec+maxtimeQ; i++ { //виртуальный таймер посекундный
		if i < allsec {
			r := rand.Intn(maxcolCustEnter)
			for j := 0; j < r; j++ {
				colCustwg.Add(1)
				go newCustomer(i, &colCustwg) //запуск покупателя в магазин

			}
			close(chsec[i])  //опопвещение
			close(chsecQ[i]) //...
			<-time.After(time.Millisecond * 10)
		} else {
			close(chsecQ[i]) //...
			<-time.After(time.Millisecond * 10)
		}
	}
	colCustwg.Wait()
	fmt.Println("Всего покупателей ", colCust)
	fmt.Printf("Done максимальное кол-во посетителей в магазине %v , текущее кол-во посетителей в магаине %v, мин. касс %v", maxcolLiveCust, colLiveCust, maxcashDesk+1)
}
