package visualizer

import (
	"github.com/oskca/gopherjs-canvas"
)

func NewRect(total int, index int, value int, ctx *canvas.Context2D) *Rectangle {
	r := new(Rectangle)
	r.Ctx = ctx
	r.Index = index
	r.ToIndex = index
	r.Left = float64((barWidth + barSpace) * index)
	r.Top = float64(canvasHeight(total) - value*heightUnit)
	r.Width = float64(barWidth)
	r.Height = float64(value * heightUnit)
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

func (r *Rectangle) Animate(timestamp float64) bool {
	finished := r.update(timestamp)
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
	r.Ctx.FillStyle = defaultColor.BarColor
	if r.IsA {
		r.Ctx.FillStyle = defaultColor.AColor
	} else if r.IsB {
		r.Ctx.FillStyle = defaultColor.BColor
	}
	r.Ctx.FillRect(r.Left, r.Top, r.Width, r.Height)
}

func (r *Rectangle) toLeft() float64 {
	return float64((barWidth + barSpace) * r.ToIndex)
}
