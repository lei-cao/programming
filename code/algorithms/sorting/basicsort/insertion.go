package basicsort

type InsertionSort struct {
	BasicSort
}

func (s *InsertionSort) Sort(a []int) {
	for i := 0; i < len(a); i++ {
		temp := a[i]
		for j := i - 1; j >= 0; j-- {
			if a[j] > temp {
				s.swap(a, j+1, j)
			} else {
				s.pass(j+1, j)
			}
		}
	}
}

