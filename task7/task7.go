package main

import (
	"fmt"
	"time"
)

type Semaphore struct {
	ch      chan struct{}
	counter int
}

func NewSemaphore(n int) *Semaphore {
	return &Semaphore{ch: make(chan struct{}, n), counter: 0}
}
func (s *Semaphore) Acquire() {
	s.ch <- struct{}{}
	s.counter++
	fmt.Println("Acquired Semaphore", s.counter)
}
func (s *Semaphore) Release() {
	<-s.ch
	s.counter--
	fmt.Println("Released Semaphore", s.counter)
}

func (s *Semaphore) Wait() {
	for {
		if s.counter == 0 {
			break
		}
	}
}

func main() {
	n := 3
	s := NewSemaphore(n)
	fmt.Println(s)
	for i := 0; i < 3; i++ {
		s.Acquire()
		go func() {
			fmt.Println("running", i)
			time.Sleep(1 * time.Second)
			fmt.Println("released", i)
			s.Release()

		}()
	}
	s.Wait()
	fmt.Println("done")
}
