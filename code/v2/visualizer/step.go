package visualizer

func NewFirstStep() Stepper {
	s := &Step{}
	s.last = s
	s.current = s
	return s
}

// Hold the operation steps queue
type Step struct {
	next    *Step
	last    *Step
	current *Step
}

func (s *Step) AddStep(stepper Stepper) {
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

func (s *Step) NextStep() Stepper {
	if s.Finished() {
		return nil
	}
	s.current = s.current.next
	return s.current
}

func (s *Step) CurrentStep() Stepper {
	return s.current
}
