package controller

import (
	"github.com/oskca/gopherjs-canvas"
	"math"
)

// Represent the element in the problem slice
type Rectangle struct {
	Index   int
	ToIndex int
	IsA     bool
	IsB     bool
	Total   int
	Value   int
	Left    float64
	Top     float64
	Width   float64
	Height  float64
	Moving  bool
	Waiting int
}

func (r *Rectangle) Draw(ctx *canvas.Context2D, delta float64) bool {
	var finished bool
	if math.Abs(r.Left-r.toLeft()) > 0.01 {
		r.Moving = true
		move := velocity * delta / 1000
		if r.Index > r.ToIndex {
			if r.Left-move < r.toLeft() {
				r.Left = r.toLeft()
			} else {
				r.Left -= move
			}
		} else if r.Index < r.ToIndex {
			if r.Left+move > r.toLeft() {
				r.Left = r.toLeft()
			} else {
				r.Left += move
			}
		}
	} else {
		r.Index = r.ToIndex
		if r.Waiting > 0 {
			r.Waiting = r.Waiting - 1
		} else {
			finished = true
		}
	}

	r.draw(ctx)
	return finished
}

func (r *Rectangle) draw(ctx *canvas.Context2D) {
	ctx.FillStyle = "#B9314F"
	if r.IsA {
		ctx.FillStyle = "#2E86AB"
	} else if r.IsB {
		//ctx.FillStyle = "#12355B"
		ctx.FillStyle = "green"
	}
	ctx.FillRect(r.Left, r.Top, r.Width, r.Height)
}

func (r *Rectangle) toLeft() float64 {
	return float64((barWidth + barSpace) * r.ToIndex)
}

func NewRect(total int, index int, value int) *Rectangle {
	r := new(Rectangle)
	r.Index = index
	r.ToIndex = index
	r.Total = total
	r.Value = value
	r.Left = float64((barWidth + barSpace) * index)
	r.Top = float64(canvasHeight(total) - value*heightUnit)
	r.Width = float64(barWidth)
	r.Height = float64(value * heightUnit)
	return r
}
