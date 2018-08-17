package ui

import (
	"github.com/oskca/gopherjs-canvas"
	"github.com/lei-cao/programming/code/v2/visualizer/defaults"
)

func NewRect(ctx *canvas.Context2D, startPoint Point, width float64, height float64, index int, value int) *Rectangle {
	r := new(Rectangle)
	r.Element = NewElement()
	r.Ctx = ctx
	r.StartPoint = startPoint
	r.DestPoint = startPoint
	r.SetWidth(width)
	r.SetHeight(height)
	r.V = value
	r.Index = index
	r.ToIndex = index
	r.OnDrawing = func() {
		if r.isA == true {
			r.Color = defaults.DefaultColor.AColor
		}
		if r.isB == true {
			r.Color = defaults.DefaultColor.BColor
		}
	}

	r.OnFinished = func() {
		r.Color = defaults.DefaultColor.BarColor
	}
	r.Color = defaults.DefaultColor.BarColor
	return r
}

// Represent the element in the problem slice
type Rectangle struct {
	*Element
	StartPoint Point
	DestPoint  Point
	Color      string
	V          int
	Index      int
	ToIndex    int
	isA        bool
	isB        bool
	OnFinished func()
	OnDrawing  func()
}

func (r *Rectangle) Animate(progress float64) bool {
	var finished bool
	r.update(progress)
	r.draw()
	if progress == 1 {
		r.StartPoint = r.DestPoint
		r.Index = r.ToIndex
		r.OnFinished()
		finished = true
	}
	return finished
}

func (r *Rectangle) update(progress float64) {
	r.StartPoint.MoveTo(r.DestPoint, progress)
}

func (r *Rectangle) draw() {
	r.OnDrawing()
	r.Ctx.FillStyle = r.Color
	r.Ctx.FillRect(r.StartPoint.X, r.StartPoint.Y, r.Width(), r.Height())
}

func (r *Rectangle) moving() bool {
	return !r.StartPoint.Equals(r.DestPoint)
}
