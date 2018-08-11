package basicsort

import (
	"github.com/lei-cao/learning-cs-again/code/visualizer"
	"github.com/lei-cao/learning-cs-again/code/algorithms/sorting"
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

type BasicSort struct {
	steps visualizer.Stepper
}

func (s *BasicSort) Steps() visualizer.Stepper {
	return s.steps
}

func (s *BasicSort) swap(nums []int, a, b int) {
	step := NewStep(a, b, true)
	s.steps.AddStep(step)
	nums[a], nums[b] = nums[b], nums[a]
}

func (s *BasicSort) pass(a, b int) {
	step := NewStep(a, b, false)
	s.steps.AddStep(step)
}

