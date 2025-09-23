package main

import (
	"fmt"
	"time"
)

// вариант 1: for - select для чтения из всех каналов
func or(channels ...<-chan interface{}) <-chan interface{} {
	done := make(chan interface{})
	go func() {
		defer close(done)
		for {
			for _, ch := range channels {
				select {
				case <-ch: // если канал закрыт -> выход -> закрытие done канала
					return
				default:
				}
			}

		}
	}()
	return done
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("done after %v", time.Since(start))
}
