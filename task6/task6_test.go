package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestMergeCh(t *testing.T) {
	in1 := []float32{1, 2, 3}
	in2 := []float32{4, 5, 6}
	in3 := []float32{7, 8, 9}

	var resOut []float32
	var resIn []float32
	resIn = append(resIn, in1...)
	resIn = append(resIn, in2...)
	resIn = append(resIn, in3...)

	c1 := make(chan float32)
	c2 := make(chan float32)
	c3 := make(chan float32)

	go func() {
		for _, n := range in1 {
			c1 <- n
		}
		close(c1)
	}()

	go func() {
		for _, n := range in2 {
			c2 <- n
		}
		close(c2)
	}()

	go func() {
		for _, n := range in3 {
			c3 <- n
		}
		close(c3)
	}()

	result := MergeCh(c1, c2, c3)
	for v := range result {
		resOut = append(resOut, v)
	}

	sort.Slice(resOut, func(i, j int) bool {
		return resOut[i] < resOut[j]
	})

	if !reflect.DeepEqual(resOut, resIn) {
		t.Errorf("MergeCh returned wrong result have %v, want %v", resOut, resIn)
	}
}
