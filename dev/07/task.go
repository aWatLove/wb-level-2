package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Or channel
	Реализовать функцию, которая будет объединять один или более
done-каналов в single-канал, если один из его составляющих каналов
закроется.
	Очевидным вариантом решения могло бы стать выражение при
использованием select, которое бы реализовывало эту связь, однако
иногда неизвестно общее число done-каналов, с которыми вы
работаете в рантайме. В этом случае удобнее использовать вызов
единственной функции, которая, приняв на вход один или более
or-каналов, реализовывала бы весь функционал.

Определение функции:

var or func(channels ...<- chan interface{}) <- chan interface{}

Пример использования функции:
sig := func(after time.Duration) <- chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}

start := time.Now()
<-or (
	sig(2*time.Hour),
	sig(5*time.Minute),
	sig(1*time.Second),
	sig(1*time.Hour),
	sig(1*time.Minute),
)

fmt.Printf(“fone after %v”, time.Since(start))
*/

func or(channels ...<-chan interface{}) <-chan interface{} {
	res := make(chan interface{})
	defer close(res)

	wg := sync.WaitGroup{}

	for _, ch := range channels {
		wg.Add(1)
		go func(ch <-chan interface{}) {
			for i := range ch {
				res <- i
			}
			wg.Done()
		}(ch)
	}

	wg.Wait()

	return res
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
		sig(15*time.Second),
		sig(1*time.Hour),
		sig(10*time.Second),
		sig(10*time.Minute),
	)

	fmt.Printf("fone after %v", time.Since(start))
}
