package main

import (
	"reflect"
	"testing"
)

func Test_addElements(t *testing.T) {
	type args struct {
		originalSlice []int
		elem          int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "good one",
			args: args{
				originalSlice: []int{1, 2, 3, 4, 5},
				elem:          10,
			},
			want: []int{1, 2, 3, 4, 5, 10},
		},
		{
			name: "another one",
			args: args{
				originalSlice: []int{1, 2, 3, 4, 5, 8},
				elem:          10,
			},
			want: []int{1, 2, 3, 4, 5, 8, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addElements(tt.args.originalSlice, tt.args.elem); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addElements() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_copySlice(t *testing.T) {
	type args struct {
		originalSlice []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test copy",
			args: args{
				originalSlice: []int{1, 2, 3, 4, 5},
			},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := copySlice(tt.args.originalSlice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("copySlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateSlice(t *testing.T) {
	s := generateSlice()
	if len(s) != 10 {
		t.Errorf("generateSlice() = %v, want 10", s)
	}
}

func Test_removeElement(t *testing.T) {
	type args struct {
		originalSlice []int
		index         int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test removeElement",
			args: args{
				originalSlice: []int{1, 2, 3, 4, 5},
				index:         1,
			},
			want:    []int{1, 3, 4, 5},
			wantErr: false,
		},
		{
			name: "test removeElement negative case",
			args: args{
				originalSlice: []int{1, 2, 3, 4, 5},
				index:         10,
			},
			want:    []int{1, 2, 3, 4, 5},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := removeElement(tt.args.originalSlice, tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("removeElement() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeElement() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sliceExample(t *testing.T) {
	type args struct {
		originalSlice []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "test Remove evens",
			args: args{
				originalSlice: []int{1, 2, 3, 4, 5},
			},
			want: []int{2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sliceExample(tt.args.originalSlice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sliceExample() = %v, want %v", got, tt.want)
			}
		})
	}
}
