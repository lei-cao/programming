package mergesort

import (
	"github.com/lei-cao/learning-cs-again/code/visualizer"
	"github.com/lei-cao/learning-cs-again/code/algorithms/sorting"
)

func NewFirstStep() visualizer.Stepper {
	s := &Step{}
	s.last = s
	s.current = s
	return s
}
func NewTopDownMergeSort() sorting.Sorter {
	m := new(TopDownMergeSort)
	m.from = "b"
	m.to = "a"
	m.steps = NewFirstStep()
	return m
}

type TopDownMergeSort struct {
	steps visualizer.Stepper
	from  string
	to    string
}

func (m *TopDownMergeSort) Steps() visualizer.Stepper {
	return m.steps
}

type IntSlice struct {
	nums []int
	name string
}

func (m *TopDownMergeSort) Sort(nums []int) {

	var numsB []int = make([]int, len(nums))
	copy(numsB, nums)
	from := &IntSlice{
		nums: nums,
		name: "a",
	}
	to := &IntSlice{
		nums: numsB,
		name: "b",
	}

	m.start(0, len(nums))
	// Sort data From numsB to nums
	m.topDownSplitMerge(to, from, 0, len(nums))
}

func (m *TopDownMergeSort) topDownSplitMerge(mergeFrom *IntSlice, mergeTo *IntSlice, iBegin int, iEnd int) {
	if iEnd-iBegin < 2 {
		return
	}
	iMid := (iBegin + iEnd) / 2
	m.split(iBegin, iEnd, mergeFrom.name, mergeTo.name)
	m.topDownSplitMerge(mergeTo, mergeFrom, iBegin, iMid)
	m.topDownSplitMerge(mergeTo, mergeFrom, iMid, iEnd)
	m.topDownMerge(mergeFrom, mergeTo, iBegin, iMid, iEnd)
}

func (m *TopDownMergeSort) topDownMerge(mergeFrom *IntSlice, mergeTo *IntSlice, iBegin int, iMid int, iEnd int) {
	i := iBegin
	j := iMid
	for k := iBegin; k < iEnd; k++ {
		if i < iMid && (j >= iEnd || mergeFrom.nums[i] <= mergeFrom.nums[j]) {
			mergeTo.nums[k] = mergeFrom.nums[i]
			m.assign(iBegin, iMid, iEnd, k, i, j, "i", mergeFrom.name, mergeTo.name)
			i += 1
		} else {
			mergeTo.nums[k] = mergeFrom.nums[j]
			m.assign(iBegin, iMid, iEnd, k, i, j, "j", mergeFrom.name, mergeTo.name)
			j += 1
		}
	}
}

func (m *TopDownMergeSort) start(iBegin, iEnd int) {
	s := &Step{}
	s.IBegin = iBegin
	s.IEnd = iEnd
	s.isFirst = true
	s.From = m.from
	s.To = m.to
	m.steps.AddStep(s)
}

func (m *TopDownMergeSort) split(iBegin, iEnd int, from, to string) {
	s := &Step{}
	s.IBegin = iBegin
	s.IEnd = iEnd
	s.isSplit = true
	s.From = from
	s.To = to
	m.steps.AddStep(s)
}

func (m *TopDownMergeSort) assign(iBegin int, iMid int, iEnd int, k int, i int, j int, assign string, from string, to string) {
	s := &Step{}
	s.IBegin = iBegin
	s.IMid = iMid
	s.IEnd = iEnd
	s.K = k
	s.I = i
	s.J = j
	s.Assign = assign
	s.From = from
	s.To = to
	s.isAssign = true
	m.steps.AddStep(s)
}


