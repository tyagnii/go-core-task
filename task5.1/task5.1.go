package main

import (
	"fmt"
	"math/rand"
)

func genRand(c chan<- int32) {
	for {
		c <- rand.Int31()
	}
}

func main() {
	c := make(chan int32)
	go genRand(c)
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
}
