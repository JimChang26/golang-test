package main

import (
	"testing"
)

func Test_add(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_double(t *testing.T) {
	type args struct {
		a int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_double", args{2}, 54},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := double(tt.args.a); got != tt.want {
				t.Errorf("double() = %v, want %v", got, tt.want)
			}
		})
	}
}
