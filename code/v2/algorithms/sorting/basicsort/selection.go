package basicsort

type SelectionSort struct {
	BasicSort
}

func (s *SelectionSort) Sort(a []int) {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				s.swap(a, i, j)
			} else {
				s.pass(i, j)
			}
		}
	}
}
