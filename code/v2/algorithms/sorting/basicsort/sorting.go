package basicsort

import (
	"github.com/lei-cao/programming/code/v2/visualizer"
	"github.com/lei-cao/programming/code/v2/algorithms/sorting"
)

func NewBubbleSort() sorting.Sorter {
	s := new(BubbleSort)
	s.steps = NewFirstStep()
	return s
}

func NewInsertionSort() sorting.Sorter {
	s := new(InsertionSort)
	s.steps = NewFirstStep()
	return s
}

func NewQuickSort() sorting.Sorter {
	s := new(QuickSort)
	s.steps = NewFirstStep()
	return s
}
func NewSelectionSort() sorting.Sorter {
	s := new(SelectionSort)
	s.steps = NewFirstStep()
	return s
}

func NewHeapSort() sorting.Sorter {
	m := new(HeapSort)
	m.steps = NewFirstStep()
	return m
}

type BasicSort struct {
	steps visualizer.Stepper
}

func (s *BasicSort) Steps() visualizer.Stepper {
	return s.steps
}

func (s *BasicSort) swap(a []int, ia, ib int) {
	step := NewStep(ia, ib, true)
	s.steps.AddStep(step)
	a[ia], a[ib] = a[ib], a[ia]
}

func (s *BasicSort) pop(a []int, ia, ib int) {
	step := NewPopStep(ia, ib, true)
	s.steps.AddStep(step)
	a[ia], a[ib] = a[ib], a[ia]
}

func (s *BasicSort) pass(a, b int) {
	step := NewStep(a, b, false)
	s.steps.AddStep(step)
}

