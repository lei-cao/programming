package sorting

type QuickSort struct {
	BasicSort
}

func (s *QuickSort) Sort(nums []int) {
	s.quickSort(nums, 0, len(nums)-1)
}

func (s *QuickSort) quickSort(nums []int, lo, hi int) {
	if lo < hi {
		p := s.partition(nums, lo, hi)
		s.quickSort(nums, lo, p-1)
		s.quickSort(nums, p+1, hi)
	}
}

func (s *QuickSort) partition(nums []int, lo, hi int) int {
	pivot := nums[hi]
	i := lo - 1
	for j := lo; j < hi; j++ {
		if nums[j] < pivot {
			i += 1
			s.swap(nums, i, j)
		} else {
			s.pass(hi, j)
		}
	}
	s.swap(nums, hi, i+1)
	return i + 1
}
