package main

import "testing"

func TestReverse(t *testing.T) {
	test(t, reverse)
}

func test(t *testing.T, f func(x int) int) {
	tables := []struct {
		x      int
		result int
	}{
		{123, 321},
		{-123, -321},
		{120, 21},
		{1534236469, 0},
	}

	for _, table := range tables {
		result := f(table.x)
		if table.result != result {
			t.Errorf("incorrect %v. Got: %v", table, result)
		}
	}
}
