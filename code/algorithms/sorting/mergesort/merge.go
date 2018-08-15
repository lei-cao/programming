package mergesort

import (
	"github.com/lei-cao/programming/code/visualizer"
	"github.com/lei-cao/programming/code/algorithms/sorting"
)

func NewFirstStep() visualizer.Stepper {
	s := &Step{}
	s.last = s
	s.current = s
	return s
}
func NewTopDownMergeSort() sorting.Sorter {
	m := new(TopDownMergeSort)
	m.steps = NewFirstStep()
	return m
}

type TopDownMergeSort struct {
	steps visualizer.Stepper
}

func (m *TopDownMergeSort) Steps() visualizer.Stepper {
	return m.steps
}

type IntSlice struct {
	a    []int
	name string
}

func (m *TopDownMergeSort) Sort(a []int) {

	var b = make([]int, len(a))
	copy(b, a)
	from := &IntSlice{
		a:    a,
		name: "a",
	}
	to := &IntSlice{
		a:    b,
		name: "b",
	}

	// Sort data From b to a
	m.split(0, len(a), to.name, from.name)
	m.topDownSplitMerge(to, from, 0, len(a))
}

func (m *TopDownMergeSort) topDownSplitMerge(mergeFrom *IntSlice, mergeTo *IntSlice, iBegin int, iEnd int) {
	if iEnd-iBegin < 2 {
		return
	}
	iMid := (iBegin + iEnd) / 2
	m.split(iBegin, iMid, mergeTo.name, mergeFrom.name)
	m.topDownSplitMerge(mergeTo, mergeFrom, iBegin, iMid)
	m.split(iMid, iEnd, mergeTo.name, mergeFrom.name)
	m.topDownSplitMerge(mergeTo, mergeFrom, iMid, iEnd)
	m.topDownMerge(mergeFrom, mergeTo, iBegin, iMid, iEnd)
}

func (m *TopDownMergeSort) topDownMerge(mergeFrom *IntSlice, mergeTo *IntSlice, iBegin int, iMid int, iEnd int) {
	i := iBegin
	j := iMid
	for k := iBegin; k < iEnd; k++ {
		if i < iMid && (j >= iEnd || mergeFrom.a[i] <= mergeFrom.a[j]) {
			mergeTo.a[k] = mergeFrom.a[i]
			m.assign(iBegin, iMid, iEnd, k, i, j, "i", mergeFrom.name, mergeTo.name)
			i += 1
		} else {
			mergeTo.a[k] = mergeFrom.a[j]
			m.assign(iBegin, iMid, iEnd, k, i, j, "j", mergeFrom.name, mergeTo.name)
			j += 1
		}
	}
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


