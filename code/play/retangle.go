package controller

import (
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
	Waiting int
}

func (r *Rectangle) Update(delta float64) bool {
	var finished bool
	if math.Abs(r.Left-r.toLeft()) > 0.01 {
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

	return finished
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
