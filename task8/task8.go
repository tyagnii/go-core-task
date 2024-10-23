package main

import (
	"fmt"
	"math"
	"math/rand"
)

func writeUint8(in chan<- uint8, numbers ...uint8) {
	for _, n := range numbers {
		in <- n
	}
	close(in)
}

func conveyor(in <-chan uint8, out chan<- float64) {
	for n := range in {
		out <- math.Pow(float64(n), 3)
	}
	close(out)
}

func main() {
	var uint8Slice []uint8
	in := make(chan uint8)
	out := make(chan float64)

	// gen slice of pseudorandom uit8
	for i := 0; i < 10; i++ {
		uint8Slice = append(uint8Slice, uint8(rand.Uint32()))
	}
	go writeUint8(in, uint8Slice...)
	go conveyor(in, out)
	for n := range out {
		fmt.Printf("%v, %e\n", n, n)

	}
}
