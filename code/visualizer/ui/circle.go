package ui

import (
	"github.com/oskca/gopherjs-canvas"
	"github.com/lei-cao/programming/code/visualizer/defaults"
	"math"
)

func NewCircle(ctx *canvas.Context2D, startPoint Point, radius float64, index int, value int) *Circle {
	r := new(Circle)
	r.Ctx = ctx
	r.StartPoint = startPoint
	r.DestPoint = startPoint
	r.Radius = radius
	r.V = value
	r.OnDrawing = func() {
		r.Color = defaults.DefaultColor.CColor
	}

	r.OnFinished = func() {
		r.Color = defaults.DefaultColor.BarColor
	}
	r.Color = defaults.DefaultColor.BarColor
	return r
}

// Represent the element in the problem slice
type Circle struct {
	Element
	StartPoint Point
	DestPoint  Point
	Radius     float64
	Color      string
	V          int
	OnFinished func()
	OnDrawing  func()
}

func (r *Circle) Animate(progress float64) bool {
	var finished bool
	r.update(progress)
	r.draw()
	if progress == 1 {
		r.StartPoint = r.DestPoint
		r.OnFinished()
		finished = true
	}
	return finished
}

func (r *Circle) update(progress float64) {
	r.StartPoint.MoveTo(r.DestPoint, progress)
}

func (r *Circle) draw() {
	r.OnDrawing()
	r.Ctx.FillStyle = r.Color
	r.Ctx.Arc(r.StartPoint.X, r.StartPoint.Y, r.Radius, 0, 2*math.Pi, true)
}

func (r *Circle) moving() bool {
	return !r.StartPoint.Equals(r.DestPoint)
}
