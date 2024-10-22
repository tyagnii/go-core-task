package main

import (
	"reflect"
	"testing"
)

func TestExclusion(t *testing.T) {
	type args struct {
		first  []string
		second []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				first:  []string{"a", "b", "c", "d", "e", "f"},
				second: []string{"a", "b", "c"},
			},
			want: []string{"d", "e", "f"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Exclusion(tt.args.first, tt.args.second); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Exclusion() = %v, want %v", got, tt.want)
			}
		})
	}
}
