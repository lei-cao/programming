package sorting

type TopDownMergeSort struct {
	BasicSort
}

func (s *TopDownMergeSort) Sort(nums []int) {
	var numsB []int = make([]int, len(nums))
	copy(numsB, nums)
	s.topDownSplitMerge(numsB, 0, len(nums), nums)
}

func (s *TopDownMergeSort) topDownSplitMerge(b []int, iBegin int, iEnd int, a []int) {
	if iEnd-iBegin < 2 {
		return
	}
	iMid := (iBegin + iEnd) / 2
	s.topDownSplitMerge(a, iBegin, iMid, b)
	s.topDownSplitMerge(a, iMid, iEnd, b)
	s.topDownMerge(b, iBegin, iMid, iEnd, a)
}

func (s *TopDownMergeSort) topDownMerge(a []int, iBegin int, iMid int, iEnd int, b []int) {
	i := iBegin
	j := iMid
	for k := iBegin; k < iEnd; k++ {
		if i < iMid && (j >= iEnd || a[i] <= a[j]) {
			b[k] = a[i]
			i += 1
			s.swap(a, k, i)
		} else {
			b[k] = a[j]
			s.swap(a, k, j)
			j += 1
		}
	}
}
