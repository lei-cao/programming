
package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSumMapBasic([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSumMap([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSumMapStruct([]int{2, 7, 11, 15}, 9))
}

// brute force O(n^2)
func twoSum(nums []int, target int) []int {
	for k, v := range nums {
		for i := k + 1; i < len(nums); i++ {
			if v+nums[i] == target {
				return []int{k, i}
			}
		}
	}
	return nil
}

// search map O(n)
func twoSumMapBasic(nums []int, target int) []int {
	searchMap := map[int]int{}
	// {[]int{1, 1, 1, 1, 0}, 1, []int{0, 4}},
	for k, v := range nums {
		searchMap[target-v] = k
	}
	// {-1:1, 0: 2, 0:3},
	for k, v := range nums {
		if searchMap[v] != 0 {
			return []int{k, searchMap[v]}
		}
	}
	return nil
}

// search map O(n)
func twoSumMap(nums []int, target int) []int {
	searchMap := map[int]int{}
	for k, v := range nums {
		if searchMap[v] != 0 {
			return []int{searchMap[v] - 1, k}
		} else {
			// Default value is 0,
			// so need to start counting index from 1
			searchMap[target-v] = k + 1
		}
	}

	return nil
}


// search map O(n)
func twoSumMapStruct(nums []int, target int) []int {
	// To explicitly set the key
	type Key struct {
		key int
		set bool
	}
	searchMap := map[int]Key{}
	for k, v := range nums {
		if searchMap[v].set {
			return []int{searchMap[v].key, k}
		} else {
			searchMap[target-v] = Key{key: k, set: true}
		}
	}

	return nil
}
