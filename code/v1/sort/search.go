package sort

import (
	"fmt"
	"time"
)

type Search interface {
	Search(val int, nums []int) int
}

func DoSearch(search Search, val int, nums []int, c chan int) {
	tmp := make([]int, len(nums))
	copy(tmp, nums)
	c <- search.Search(val, tmp)
}

type DefaultSearch struct {
}

func (s *DefaultSearch) Search(val int, nums []int) int {
	obj := createCanvas("Search", len(nums))
	fmt.Println(val)

	draw(nums, 0, -1, obj)
	for k, v := range nums {
		draw(nums, k, -1, obj)
		if v == val {
			return k
		}
		time.Sleep(100 * time.Millisecond)
	}
	return -1
}
