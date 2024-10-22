package main

import (
	"reflect"
	"testing"
)

func Test_genRand(t *testing.T) {
	type args struct {
		c chan int32
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				c: make(chan int32),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go genRand(tt.args.c)
			tmp := <-tt.args.c
			if reflect.TypeOf(tmp).Name() != "int32" {
				t.Error("wrong type")

			}
		})
	}
}
