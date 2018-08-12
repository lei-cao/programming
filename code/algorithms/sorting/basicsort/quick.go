package basicsort

type QuickSort struct {
	BasicSort
}

func (s *QuickSort) Sort(a []int) {
	s.quickSort(a, 0, len(a)-1)
}

func (s *QuickSort) quickSort(a []int, lo, hi int) {
	if lo < hi {
		p := s.partition(a, lo, hi)
		s.quickSort(a, lo, p-1)
		s.quickSort(a, p+1, hi)
	}
}

func (s *QuickSort) partition(a []int, lo, hi int) int {
	pivot := a[hi]
	i := lo - 1
	for j := lo; j < hi; j++ {
		if a[j] < pivot {
			i += 1
			s.swap(a, i, j)
		} else {
			s.pass(hi, j)
		}
	}
	s.swap(a, hi, i+1)
	return i + 1
}
