package main

import (
	"fmt"
	"time"
)

func or(channels ...<-chan interface{}) <-chan interface{} {
	// Создаем канал, который будет возвращен как результат
	result := make(chan interface{})

	// Запускаем отдельную горутину для каждого входящего канала
	for _, ch := range channels {
		go func(ch <-chan interface{}) {
			for {
				select {
				// Когда один из входящих каналов закрывается, закрываем и результирующий канал
				case <-ch:
					close(result)
					return
				}
			}
		}(ch)
	}

	return result
}

func sig(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}

func main() {
	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("Done after %v", time.Since(start))
}
