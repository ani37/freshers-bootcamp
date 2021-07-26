package main

import (
	"fmt"
	"testing"
)


func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestBalance(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {

		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

func TestBalanceDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{50, 60, 490},
		{100, 100, 500},
		{50, 550, 0},
		{0, 0, 500},
		{-1, 500, 0},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			withdrawal(tt.a)
			deposit(tt.b)
			ans := balance
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}