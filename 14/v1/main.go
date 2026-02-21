package main

import (
	"fmt"
	"time"
)

// рекурсивно объединять каналы
func or(channels ...<-chan interface{}) <-chan interface{} {
	if len(channels) == 0 {
		return nil
	}
	if len(channels) == 1 {
		return channels[0]
	}
	done := make(chan interface{})
	go func() {
		defer close(done)
		for {
			select {
			case <-channels[0]:
				return
			case <-or(channels[1:]...):
				return
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
