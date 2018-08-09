package visualizer

type Screener interface {
	Ready() bool
	Update(step Stepper)
	Swap(ia, ib int)
	Pass(ia, ib int)
	Draw(timestamp float64)
	Clear()
}

type Stepper interface {
	Swapper
	AddStep(a, b int, doSwap bool)
	Finished() bool
	NextStep() Stepper
	CurrentStep() Stepper
}

type Swapper interface {
	A() int
	B() int
	DoSwap() bool
}