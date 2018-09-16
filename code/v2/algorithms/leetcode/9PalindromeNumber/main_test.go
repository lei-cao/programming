package main

import "testing"

func TestReverse(t *testing.T) {
	test(t, isPalindrome)
}

func test(t *testing.T, f func(x int) bool) {
	tables := []struct {
		x      int
		result bool
	}{
		{121, true},
		{-121, false},
		{10, false},
	}

	for _, table := range tables {
		result := f(table.x)
		if table.result != result {
			t.Errorf("incorrect %v. Got: %v", table, result)
		}
	}
}
