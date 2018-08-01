package sort

import "time"

type SelectionSort struct {
}

func (s *SelectionSort) Sort(nums []int, id string) {
	obj := createCanvas(id, len(nums))

	draw(nums, 0, 0, obj)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
			draw(nums, i, j, obj)
			time.Sleep(1 * time.Millisecond)
		}
	}
}
