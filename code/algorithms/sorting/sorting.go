package sorting

import "github.com/lei-cao/learning-cs-again/code/visualizer"

type Sorter interface {
	Sort(nums []int)
	Steps() visualizer.Stepper
}

func NewBubbleSort() Sorter {
	s := new(BubbleSort)
	s.steps = visualizer.NewStep()
	return s
}

func NewInsertionSort() Sorter {
	s := new(InsertionSort)
	s.steps = visualizer.NewStep()
	return s
}
func NewTopDownMergeSort() Sorter {
	s := new(TopDownMergeSort)
	s.steps = visualizer.NewStep()
	return s
}
func NewQuickSort() Sorter {
	s := new(QuickSort)
	s.steps = visualizer.NewStep()
	return s
}
func NewSelectionSort() Sorter {
	s := new(SelectionSort)
	s.steps = visualizer.NewStep()
	return s
}

type BasicSort struct {
	steps visualizer.Stepper
}

func (s *BasicSort) swap(nums []int, a, b int) {
	s.steps.AddStep(a, b, true)
	nums[a], nums[b] = nums[b], nums[a]
}

func (s *BasicSort) pass(a, b int) {
	s.steps.AddStep(a, b, false)
}

func (s *BasicSort) Steps() visualizer.Stepper {
	return s.steps
}
