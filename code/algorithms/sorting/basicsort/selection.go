package basicsort

type SelectionSort struct {
	BasicSort
}

func (s *SelectionSort) Sort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				s.swap(nums, i, j)
			} else {
				s.pass(i, j)
			}
		}
	}
}
