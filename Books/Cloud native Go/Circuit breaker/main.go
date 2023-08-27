package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

type Circuit func(context.Context) (string, error)

func Breaker(circuit Circuit, rejectionCount uint) Circuit {
	var consecutiveFailures int = 0
	var lastAttempt = time.Now()
	var m sync.RWMutex

	return func(ctx context.Context) (string, error) {
		m.RLock() // Установить "блокировку чтения"
		d := consecutiveFailures - int(rejectionCount)
		if d >= 0 {
			shouldRetryAt := lastAttempt.Add(time.Second * time.Duration(d)) //(d))
			if !time.Now().After(shouldRetryAt) {
				m.RUnlock()
				return "", errors.New("service unreachable")
			}
		}
		m.RUnlock()                   // Освободить блокировку чтения
		response, err := circuit(ctx) // Послать запрос, как обычно
		m.Lock()                      // Заблокировать общие ресурсы
		defer m.Unlock()
		lastAttempt = time.Now() // Зафиксировать время попытки
		if err != nil {          // Если Circuit вернула ошибку,
			consecutiveFailures++ // увеличить счетчик ошибок
			return response, err  // и вернуть ошибку
		}
		consecutiveFailures = 0 // Сбросить счетчик ошибок
		return response, nil
	}
}

func main() {

	fn := func(ctx context.Context) (string, error) {

		sec := time.Now().Second()
		fmt.Println(sec)
		if sec%9 != 0 {
			return "error", errors.New("one more time")
		}
		return "done", nil
	}
	var ctx context.Context
	br := Breaker(fn, 3)

	str, err := br(ctx)

	for err != nil {

		time.Sleep(1000 * time.Millisecond)
		fmt.Println(err)

		str, err = br(ctx)
	}

	fmt.Println(str)
}
