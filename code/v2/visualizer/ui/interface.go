package ui

import "github.com/lei-cao/programming/code/v2/visualizer"

type Elementer interface {
	Width() float64
	Height() float64
	Update(stepper visualizer.Stepper)
	Draw(progress float64)
	Ready() bool
}

const (
	EASING_IN = iota
	EASING_OUT
)

type Easinger interface {
	Name(int)
	Timing(progress float64) float64
}

type Transitioner interface {
	Duration(duration float64)
	Easing(name int)
	From() float64
	To() float64
	Value(timeDelta float64) float64
	OnTransitionStart()
	OnTransitionEnd()
}
