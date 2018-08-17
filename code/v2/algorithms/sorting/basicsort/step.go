package basicsort

import "github.com/lei-cao/programming/code/v2/visualizer"

func NewFirstStep() visualizer.Stepper {
	s := &Step{}
	s.last = s
	s.current = s
	return s
}

func NewStep(a int, b int, doSwap bool) visualizer.Stepper {
	s := &Step{}
	s.a = a
	s.b = b
	s.doSwap = doSwap
	return s
}

// Hold the operation steps queue
type Step struct {
	a       int
	b       int
	doSwap  bool
	next    *Step
	last    *Step
	current *Step
}

func (s *Step) AddStep(stepper visualizer.Stepper) {
	if step, ok := stepper.(*Step); ok {
		s.last.next = step
		s.last = step
	} else {
		println("Can't add step")
	}
}

func (s *Step) Finished() bool {
	return s.current.next == nil
}

func (s *Step) NextStep() visualizer.Stepper {
	if s.Finished() {
		return nil
	}
	s.current = s.current.next
	return s.current
}

func (s *Step) CurrentStep() visualizer.Stepper {
	return s.current
}

func (s *Step) A() int {
	return s.a
}
func (s *Step) B() int {
	return s.b
}
func (s *Step) DoSwap() bool {
	return s.doSwap
}
