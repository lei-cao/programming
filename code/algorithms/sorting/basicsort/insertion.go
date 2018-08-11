package basicsort

type InsertionSort struct {
	BasicSort
}

func (s *InsertionSort) Sort(nums []int) {
	for i := 0; i < len(nums); i++ {
		temp := nums[i]
		for j := i - 1; j >= 0; j-- {
			if nums[j] > temp {
				s.swap(nums, j+1, j)
			} else {
				s.pass(j+1, j)
			}
		}
	}
}

