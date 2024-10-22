package main

import (
	"reflect"
	"testing"
)

func TestIntersection(t *testing.T) {
	type args struct {
		first  []int
		second []int
	}
	tests := []struct {
		name  string
		args  args
		want  []int
		want1 bool
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				first:  []int{1, 2, 3},
				second: []int{1, 2, 4, 5, 6},
			},
			want:  []int{1, 2},
			want1: true,
		},
		{
			name: "2",
			args: args{
				first:  []int{1, 2, 3},
				second: []int{4, 5, 6},
			},
			want:  []int{},
			want1: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Intersection(tt.args.first, tt.args.second)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intersection() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Intersection() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
