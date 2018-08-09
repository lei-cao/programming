package visualizer

import (
	"github.com/oskca/gopherjs-canvas"
)

var BarWidth = 8
var BarSpace = 2
var HeightUnit = 5

func NewRect(total int, index int, value int, ctx *canvas.Context2D) *Rectangle {
	r := new(Rectangle)
	r.Ctx = ctx
	r.Index = index
	r.ToIndex = index
	r.Left = float64((BarWidth + BarSpace) * index)
	r.Top = float64(maxHeight(total) - value*HeightUnit)
	r.Width = float64(BarWidth)
	r.Height = float64(value * HeightUnit)
	return r
}

// Represent the element in the problem slice
type Rectangle struct {
	Ctx     *canvas.Context2D
	Index   int
	ToIndex int
	IsA     bool
	IsB     bool
	Left    float64
	Top     float64
	Width   float64
	Height  float64
}

func (r *Rectangle) Animate(progress float64) bool {
	finished := r.update(progress)
	r.draw()
	return finished
}

func (r *Rectangle) update(progress float64) bool {
	var finished bool

	if r.Left-r.toLeft() > 0 {
		// Moving Left
		r.Left -= (r.Left - r.toLeft()) * progress
	} else if r.Left-r.toLeft() < -0 {
		// Moving Right
		r.Left += (r.toLeft() - r.Left) * progress
	} else {
		// Stand
		r.Index = r.ToIndex
	}
	if progress == 1 {
		finished = true
	}
	return finished
}

func (r *Rectangle) draw() {
	r.Ctx.FillStyle = DefaultColor.BarColor
	if r.IsA {
		r.Ctx.FillStyle = DefaultColor.AColor
	} else if r.IsB {
		r.Ctx.FillStyle = DefaultColor.BColor
	}
	r.Ctx.FillRect(r.Left, r.Top, r.Width, r.Height)
}

func (r *Rectangle) toLeft() float64 {
	return float64((BarWidth + BarSpace) * r.ToIndex)
}

func maxHeight(size int) int {
	return size * HeightUnit
}