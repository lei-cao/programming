package basicsort

type BubbleSort struct {
	BasicSort
}

func (s *BubbleSort) Sort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-1; j++ {
			if nums[j] > nums[j+1] {
				s.swap(nums, j, j+1)
			} else {
				s.pass(j, j+1)
			}
		}
	}
}
