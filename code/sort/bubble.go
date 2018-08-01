package sort

import "time"

type BubbleSort struct {
}

func (s *BubbleSort) Sort(nums []int, id string) {
	obj := createCanvas(id, len(nums))

	draw(nums, 0, 0, obj)
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
			draw(nums, i, j, obj)
			time.Sleep(1 * time.Millisecond)
		}
	}
}

type BubbleSortSwapped struct {
}

func (s *BubbleSortSwapped) Sort(nums []int, id string) {
	obj := createCanvas(id, len(nums))

	draw(nums, 0, 0, obj)
	for i := 0; i < len(nums); i++ {
		swapped := false
		for j := 0; j < len(nums)-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				swapped = true
			}
			draw(nums, i, j, obj)
			time.Sleep(1 * time.Millisecond)
		}
		if !swapped {
			break
		}
	}
}
