package main

import "testing"

func TestTwoSum(t *testing.T) {
	test(t, twoSum)
}

func TestTwoSumMapBasic(t *testing.T) {
	test(t, twoSumMapBasic)
}
func TestTwoSumMap(t *testing.T) {
	test(t, twoSumMap)
}
func TestTwoSumMapStruct(t *testing.T) {
	test(t, twoSumMapStruct)
}

func test(t *testing.T, f func([]int, int) []int) {
	tables := []struct {
		nums   []int
		target int
		result []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{1, 2, 3}, 3, []int{0, 1}},
		{[]int{10, 20, 1}, 11, []int{0, 2}},
		{[]int{1, 0, 0}, 0, []int{1, 2}},
		{[]int{0, 1, 0}, 0, []int{0, 2}},
	}

	for _, table := range tables {
		result := f(table.nums, table.target)
		if len(table.result) != len(result) {
			t.Errorf("incorrect %v. Got: %v", table, result)
		}
		for k, v := range result {
			if table.result[k] != v {
				t.Errorf("incorrect %v. Got: %v", table, result)
			}
		}
	}
}
