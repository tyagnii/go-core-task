package main

import (
	"fmt"
	"sync"
)

func MergeCh[T any](channels ...<-chan T) chan T {
	var wg sync.WaitGroup
	output := make(chan T)

	worker := func(ch <-chan T) {
		for n := range ch {
			output <- n
		}
		wg.Done()
	}

	wg.Add(len(channels))
	for _, c := range channels {
		go worker(c)
	}

	go func() {
		wg.Wait()
		close(output)
	}()

	return output

}

func main() {
	c1 := make(chan float32)
	c2 := make(chan float32)
	c3 := make(chan float32)

	go func() {
		for _, n := range []float32{1, 3, 5} {
			c1 <- n
		}
		close(c1)
	}()

	go func() {
		for _, n := range []float32{1.34, 3.24, 8.5} {
			c2 <- n
		}
		close(c2)
	}()

	go func() {
		for _, n := range []float32{1.34e10, 3.24e20, 8.5e10} {
			c3 <- n
		}
		close(c3)
	}()

	result := MergeCh(c1, c2, c3)
	for v := range result {
		fmt.Println(v)
	}

}
