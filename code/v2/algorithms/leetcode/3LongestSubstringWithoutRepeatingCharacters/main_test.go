package main

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	test(t, lengthOfLongestSubstring)
	test(t, lengthOfLongestSubstringMap)
}

func test(t *testing.T, f func(s string) int) {
	tables := []struct {
		s string
		result int
	}{
		{"abba", 2},
		{"dvdf", 3},
		{"abcabcbb", 3},
		{"bbbb", 1},
		{"pwwkew", 3},
		{" ", 1},
		{"a", 1},
		{"aa", 1},
		{"ab", 2},
	}

	for _, table := range tables {
		result := f(table.s)
		if table.result != result {
			t.Errorf("incorrect %v. Got: %v", table, result)
		}
	}
}
