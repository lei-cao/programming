package main

import "testing"

func TestReverse(t *testing.T) {
	test(t, longestCommonPrefix)
}

func test(t *testing.T, f func(strs []string) string) {
	tables := []struct {
		strs      []string
		result string
	}{
		{[]string{}, ""},
		{[]string{"aa","a"}, "a"},
		{[]string{"flower","flow","flight",}, "fl"},
		{[]string{"dog","racecar","car"}, ""},
	}

	for _, table := range tables {
		result := f(table.strs)
		if table.result != result {
			t.Errorf("incorrect %v. Got: %v", table, result)
		}
	}
}
