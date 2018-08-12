package basicsort

type BubbleSort struct {
	BasicSort
}

func (s *BubbleSort) Sort(a []int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-1; j++ {
			if a[j] > a[j+1] {
				s.swap(a, j, j+1)
			} else {
				s.pass(j, j+1)
			}
		}
	}
}
