package tdd

import (
	"testing"
)

func TestAdd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1 + 1 = 2",
			args: args{a: 1, b: 1},
			want: 2,
		},
		{
			name: "1 + -2 = 1",
			args: args{a: 1, b: -2},
			want: 1,
		},
		{
			name: "name 1 + -1 = 1 при 0 возвращать 1",
			args: args{a: 1, b: -1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func AddMock(a int, b int) int {
	if a == 1 && b == 1 {
		return 2
	}
	if a == 1 && b == -2 {
		return 1
	}
	return 0
}

func TestSummattor(t *testing.T) {
	type args struct {
		a int
		b int
		f func(int, int) int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1 + 1 = 2",
			args: args{a: 1, b: 1, f: AddMock},
			want: 2,
		},
		{
			name: "1 + -2 = -1",
			args: args{a: 1, b: -2, f: AddMock},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Summattor(tt.args.a, tt.args.b, tt.args.f); got != tt.want {
				t.Errorf("Summattor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddSummattor(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1 + 1 = 2",
			args: args{a: 1, b: 1},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddSummattor(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("AddSummattor() = %v, want %v", got, tt.want)
			}
		})
	}
}
