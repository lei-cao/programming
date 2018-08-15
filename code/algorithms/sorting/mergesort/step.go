package mergesort

import "github.com/lei-cao/programming/code/visualizer"

// Hold the operation steps queue
type Step struct {
	IBegin   int
	IMid     int
	IEnd     int
	K        int
	I        int
	J        int
	From     string
	To       string
	Assign   string
	isFirst  bool
	isSplit  bool
	isAssign bool
	next     *Step
	last     *Step
	current  *Step
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

func (s *Step) IsFirstStep() bool {
	return s.isFirst
}

func (s *Step) IsSplitStep() bool {
	return s.isSplit
}

func (s *Step) IsAssignStep() bool {
	return s.isAssign
}