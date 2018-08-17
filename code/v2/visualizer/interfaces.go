package visualizer

type Screener interface {
	Ready() bool
	Update(step Stepper)
	Draw(progress float64)
	Clear()
}

type Stepper interface {
	AddStep(step Stepper)
	Finished() bool
	NextStep() Stepper
	CurrentStep() Stepper
}

type Swapper interface {
	A() int
	B() int
	DoSwap() bool
}
