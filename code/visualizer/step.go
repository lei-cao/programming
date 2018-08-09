package visualizer

func NewStep() Stepper {
	s := &Step{}
	s.last = s
	s.current = s
	return s
}

// Hold the operation steps queue
type Step struct {
	Swapper
	a       int
	b       int
	doSwap  bool
	next    *Step
	last    *Step
	current *Step
}

func (s *Step) AddStep(a, b int, doSwap bool) {
	step := &Step{a: a, b: b, doSwap: doSwap}
	s.last.next = step
	s.last = step
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

func (s *Step) A() int {
	return s.a
}
func (s *Step) B() int {
	return s.b
}
func (s *Step) DoSwap() bool {
	return s.doSwap
}
