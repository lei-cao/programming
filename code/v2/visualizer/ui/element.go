package ui

import (
	"github.com/oskca/gopherjs-canvas"
	"github.com/lei-cao/programming/code/v2/visualizer"
)

func NewElement() *Element {
	e := &Element{}
	e.Children = make([]Elementer, 0)
	return e
}

// The Base UI Element
type Element struct {
	Id               string
	Ctx              *canvas.Context2D
	Children         []Elementer
	AutoWidth        bool
	AutoHeight       bool
	OnFinished       func()
	OnDrawing        func()
	width            float64
	height           float64
	calculatedWidth  float64
	calculatedHeight float64
}

func (e *Element) Width() float64 {
	return e.calculatedWidth
}

func (e *Element) Height() float64 {
	return e.calculatedHeight
}
func (e *Element) Update(stepper visualizer.Stepper) {
	for k := range e.Children {
		e.Children[k].Update(stepper)
	}
}
func (e *Element) Draw(progress float64) {
	for k := range e.Children {
		e.Children[k].Draw(progress)
	}
}
func (e *Element) Ready() bool {
	for k := range e.Children {
		if !e.Children[k].Ready() {
			return false
		}
	}
	return true
}

func (e *Element) SetWidth(width float64) {
	e.width = width
	e.Resize()
}

func (e *Element) SetHeight(height float64) {
	e.height = height
	e.Resize()
}

func (e *Element) Resize() {
	var w float64
	var h float64
	for k := range e.Children {
		w += e.Children[k].Width()
		h += e.Children[k].Height()
	}
	e.calculatedWidth = e.width + w
	e.calculatedHeight = e.height + h
}
